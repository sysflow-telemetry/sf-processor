//
// Copyright (C) 2021 IBM Corporation.
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
	"io"
	"os"
	"path/filepath"
	"plugin"
	"time"

	cache "github.com/patrickmn/go-cache"
	"golang.org/x/net/context"
	"github.com/docker/docker/client"

	"github.com/sysflow-telemetry/sf-apis/go/config"
	"github.com/sysflow-telemetry/sf-apis/go/container/storage"
	"github.com/sysflow-telemetry/sf-apis/go/ioutils"
	"github.com/sysflow-telemetry/sf-apis/go/logger"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)


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
func (ah *ActionHandler) loadUserActions(dir string) {
	ah.UserDefinedActions = make(ActionMap)
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
			registerAction(ah.UserDefinedActions, name, action.GetFunc())
		}
	}
}

type ActionHandler struct {
	cstoreCli   *storage.CStorage
	dockerCli   *client.Client

	// Map of registered actions
	BuiltInActions ActionMap
	UserDefinedActions ActionMap

	maxFileSize int64
	hashTable   *cache.Cache
}

func NewActionHandler(conf Config) *ActionHandler{
	ah := new(ActionHandler)

	// Create container cli
	if conf.ContainerConfig != sfgo.Zeros.String {
		vconf, err := config.GetConfig(conf.ContainerConfig)
		if err != nil {
			logger.Error.Printf("Failed to load container config file: %s", err.Error())
		} else {
			ah.cstoreCli, err = storage.NewContainerStore(vconf)
			if err != nil {
				logger.Error.Printf("Failed to initialize storage client: %s", err.Error())
			}
		}
	}

	// Create docker cli
	var err error
	ah.dockerCli, err = client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		logger.Error.Printf("Failed to initialize docker client %s", err.Error())
	}

	// Set max hash file size
	if conf.MaxFileSize != sfgo.Zeros.Int64 {
		ah.maxFileSize = conf.MaxFileSize
	} else {
		ah.maxFileSize = MAX_FILE_SIZE
	}

	// Create hash cache
	ah.hashTable = cache.New(time.Duration(conf.HashCacheExp)*time.Minute, time.Duration(conf.HashCachePurge)*time.Minute)

	// Register built-in actions
	ah.registerBuiltIns()

	// Load user-defined actions
	ah.loadUserActions(conf.ActionDir)

	return ah
}

// HandleAction handles actions defined in rule.
func (ah *ActionHandler) HandleActions(rule Rule, r *Record) {
	for _, a := range rule.Actions {
		action, ok := ah.BuiltInActions[a]
		if !ok {
			action, ok = ah.UserDefinedActions[a]
		}
		if !ok {
			logger.Error.Println("Unknown action: '" + a + "'")
			continue
		}

		if err := action(r); err != nil {
			logger.Error.Println("Error in action: " + err.Error())
		}
	}
}

// Names of built-in actions
const (
	ACTION_HASH_PROC = "hash_proc"
	ACTION_HASH_FILE = "hash_file"
)

// Registers built-in actions
func (ah *ActionHandler) registerBuiltIns() {
	ah.BuiltInActions = make(ActionMap)
	registerAction(ah.BuiltInActions, ACTION_HASH_PROC, ah.HashProcFunc)
	registerAction(ah.BuiltInActions, ACTION_HASH_FILE, ah.HashFileFunc)
}

// Built-in hash actions 

func (ah *ActionHandler) HashProcFunc(r *Record) error {
	hs, err := ah.getHashes(r, SF_PROC_EXE)
	if err == nil {
		r.Ctx.SetHashes(HASH_TYPE_PROC, hs)
	}
	return err
}

func (ah *ActionHandler) HashFileFunc(r *Record) error {
	if Mapper.MapStr(SF_FILE_TYPE)(r) != "f" || Mapper.MapStr(SF_FILE_PATH)(r) == sfgo.Zeros.String {
		return nil
	}

	hs, err := ah.getHashes(r, SF_FILE_PATH)
	if err == nil {
		r.Ctx.SetHashes(HASH_TYPE_FILE, hs) 
	}
	return err
}

// Utility functions for hash actions

func (ah *ActionHandler) getHashes(r *Record, srcFd string) (hs *HashSet, err error) {
	contId := Mapper.MapStr(SF_CONTAINER_ID)(r)
	fPath  := Mapper.MapStr(srcFd)(r)

	key := contId + filepath.Clean(fPath)
	if entry, ok := ah.hashTable.Get(key); ok {
		hs = entry.(*HashSet)
		return
	}

	if contId == sfgo.Zeros.String {
		hs, err = ah.getHashesFromLocal(fPath)
	} else {
		if Mapper.MapStr(SF_CONTAINER_TYPE)(r) == "DOCKER" {
			hs, err = ah.getHashesFromDocker(contId, fPath)
		} else {
			hs, err = ah.getHashesFromCStore(contId, fPath)
		}
	}

	if err == nil {
		ah.hashTable.Set(key, hs, cache.DefaultExpiration)
	}
	return
}

func (ah *ActionHandler) getHashesFromLocal(path string) (*HashSet, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if fi.Size() >= ah.maxFileSize {
		return nil, errors.New("file size exceeds hashing max")
	}

	return computeHashes(file)
}

func (ah *ActionHandler) getHashesFromCStore(contId string, filePath string) (*HashSet, error) {
	if ah.cstoreCli == nil {
		return nil, errors.New("Storage client not configured")
	}

	path, _, err := ah.cstoreCli.GetContainerFilePath(contId, filePath)
	if err != nil {
		return nil, err
	}

	return ah.getHashesFromLocal(path)
}

const BUFFER_SIZE = 1024

func (ah *ActionHandler) getHashesFromDocker(contId string, filePath string) (*HashSet, error) {
	if ah.dockerCli == nil {
		return nil, errors.New("Docker client not configured")
	}

	tarStream, _, err := ah.dockerCli.CopyFromContainer(context.Background(), contId, filePath)
	if err != nil {
		return nil, err
	}
	defer tarStream.Close()

	tr := tar.NewReader(tarStream) 
	th, err := tr.Next()
	if err != nil {
		return nil, err
	}
	if th.Size >= ah.maxFileSize {
		return nil, errors.New("file size exceeds hashing max")
	}

	return computeHashes(tr)
}

func computeHashes(rd io.Reader) (hs *HashSet, err error) {
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
			err = nil // Don't treat EOF as an error
			break
		}
	}
	hs = &HashSet{
		Md5: hex.EncodeToString(m5.Sum(nil)),
		Sha1: hex.EncodeToString(s1.Sum(nil)),
		Sha256: hex.EncodeToString(s256.Sum(nil)),
	}
	return 
}

