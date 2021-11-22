//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Andreas Schade <san@zurich.ibm.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package engine implements a rules engine for telemetry records.
package engine

import (
	"archive/tar"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"plugin"

	"golang.org/x/net/context"
	"github.com/docker/docker/client"

	"github.com/sysflow-telemetry/sf-apis/go/config"
	"github.com/sysflow-telemetry/sf-apis/go/container/agents"
	"github.com/sysflow-telemetry/sf-apis/go/ioutils"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)

const CONTAINER_CONFIG = "CONTAINER_CONFIG"

var hashAgent *agents.HashAgent
var dockerCli *client.Client

func init() {
	confPath := os.Getenv(CONTAINER_CONFIG)
	if confPath != sfgo.Zeros.String {
		vconf, err := config.GetConfig(confPath)
		if err != nil {
			fmt.Printf("FATAL: Failed to load container config file: %s\n", err.Error())
		} else {
			hashAgent, err = agents.NewHashAgent(vconf)
			if err != nil {
				fmt.Printf("FATAL: Failed to initialize hash agent: %s\n", err.Error())
			}
		}
	}

	var err error
	dockerCli, err = client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		fmt.Printf("FATAL: Failed to initialize docker client %s\n", err.Error())
	}
}

// Prototype of an action function
type ActionFunc func(r *Record) error

type ActionMap map[string]ActionFunc

// Action interface for user-defined actions
type Action interface {
	GetName() string
	GetFunc() ActionFunc
}

const ActionSym = "Action"

// Registers an action function
func registerAction(reg ActionMap, name string, f ActionFunc) {
        if  _, ok := reg[name]; ok {
                logger.Warn.Println("Re-declaration of action '" + name + "'")
        }
        reg[name] = f
}

// LoadActions loads user-defined actions from path
func (pi *PolicyInterpreter) loadUserActions(dir string) {
	pi.userDefinedActions = make(ActionMap)
	if paths, err := ioutils.ListFilePaths(dir, ".so"); err == nil {
		var plug *plugin.Plugin
		for _, path := range paths {
			logger.Info.Println("Loading user-defined actions from file " + path)
			if plug, err = plugin.Open(path); err != nil {
				logger.Error.Println(err.Error())
				continue
			}
			sym, err := plug.Lookup(ActionSym)
			if err != nil {
				logger.Error.Println(err.Error())
				continue
			}
			action, ok := sym.(Action)
			if !ok {
				logger.Error.Println("Action symbol loaded from " + path + " must implement Action interface")
				continue
			}

			name := action.GetName()
			logger.Info.Println("Registering user-defined action '" + name + "'")
			registerAction(pi.userDefinedActions, name, action.GetFunc())
		}
	}
}

// Names of built-in actions
const (
	ACTION_HASH_PROC = "hash_proc"
	ACTION_HASH_FILE = "hash_file"
)

// Registers built-in actions
func (pi *PolicyInterpreter) registerBuiltIns() {
	pi.builtInActions = make(ActionMap)
	registerAction(pi.builtInActions, ACTION_HASH_PROC, HashProcFunc)
	registerAction(pi.builtInActions, ACTION_HASH_FILE, HashFileFunc)
}

// Built-in hash actions 

func HashProcFunc(r *Record) error {
	m1s, s1s, s256s, err := getHashes(r, SF_PROC_EXE)
	if err == nil {
		r.Ctx.SetHashes(HASH_TYPE_PROC, m1s, s1s, s256s)
	}
	return err
}

func HashFileFunc(r *Record) error {
	if Mapper.MapStr(SF_FILE_TYPE)(r) != "f" || Mapper.MapStr(SF_FILE_PATH)(r) == sfgo.Zeros.String {
		return nil
	}

	m1s, s1s, s256s, err := getHashes(r, SF_FILE_PATH)
	if err == nil {
		r.Ctx.SetHashes(HASH_TYPE_FILE, m1s, s1s, s256s) 
	}
	return err
}

// Utility functions for hash actions

func getHashes(r *Record, srcFd string) (m5s string, s1s string, s256s string, err error) {
	contId := Mapper.MapStr(SF_CONTAINER_ID)(r)
	filePath := Mapper.MapStr(srcFd)(r)

	if contId == sfgo.Zeros.String {
		m5s, s1s, s256s, err = getHashesFromLocal(filePath)
	} else {
		if Mapper.MapStr(SF_CONTAINER_TYPE)(r) == "DOCKER" {
			m5s, s1s, s256s, err = getHashesFromDocker(contId, filePath)
		} else {
			if hashAgent == nil {
				err = errors.New("hash agent not initialized")
				return
			}
			m5s, s1s, s256s, _, _, err = hashAgent.GetHashes(contId, filePath)
		}
	}
	return
}

func getHashesFromLocal(path string) (m5s string, s1s string, s256s string, err error) {
	var file *os.File
	file, err = os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	return computeHashes(file)
}

const BUFFER_SIZE = 1024

func getHashesFromDocker(contId string, filePath string) (m5s string, s1s string, s256s string, err error) {
	if dockerCli == nil {
		err = errors.New("docker client not initialized")
		return
	}
	tarStream, _, err := dockerCli.CopyFromContainer(context.Background(), contId, filePath)
	if err != nil {
		return
	}
	defer tarStream.Close()
	tr := tar.NewReader(tarStream)
	if _, err = tr.Next(); err != nil {
		return
        }

	return computeHashes(tr)
}

func computeHashes(rd io.Reader) (m5s string, s1s string, s256s string, err error) {
	m5 := md5.New()
	s1 := sha1.New()
	s256 := sha256.New()

	bytesread := 0
	buffer := make([]byte, BUFFER_SIZE)
        for {
                bytesread, err = rd.Read(buffer)
                if err != nil && err != io.EOF {
			return
                }
                if bytesread > 0 {
                        m5.Write(buffer[:bytesread])
                        s1.Write(buffer[:bytesread])
                        s256.Write(buffer[:bytesread])
                }
                if err == io.EOF {
                        break
                }
        }
        m5s = hex.EncodeToString(m5.Sum(nil))
	s1s = hex.EncodeToString(s1.Sum(nil))
	s256s = hex.EncodeToString(s256.Sum(nil))
	err = nil
	return 
}
