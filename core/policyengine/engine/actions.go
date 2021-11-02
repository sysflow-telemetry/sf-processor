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
	"encoding/hex"
	"crypto"
	"io"
	"os"
	"plugin"

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
	HASH_MD5_PROC = "hash_md5_proc"
	HASH_MD5_FILE = "hash_md5_file"
	HASH_SHA1_PROC = "hash_sha1_proc"
	HASH_SHA1_FILE = "hash_sha1_file"
	HASH_SHA256_PROC = "hash_sha256_proc"
	HASH_SHA256_FILE = "hash_sha256_file"
)

// Registers built-in actions
func (pi *PolicyInterpreter) registerBuiltIns() {
	pi.builtInActions = make(ActionMap)
	registerAction(pi.builtInActions, HASH_MD5_PROC, HashMd5ProcFunc)
	registerAction(pi.builtInActions, HASH_MD5_FILE, HashMd5FileFunc)
	registerAction(pi.builtInActions, HASH_SHA1_PROC, HashSha1ProcFunc)
	registerAction(pi.builtInActions, HASH_SHA1_FILE, HashSha1FileFunc)
	registerAction(pi.builtInActions, HASH_SHA256_PROC, HashSha256ProcFunc)
	registerAction(pi.builtInActions, HASH_SHA256_FILE, HashSha256FileFunc)
}

// Built-in hash actions 

func HashMd5ProcFunc(r *Record) error {
	hs, err := computeHash(crypto.MD5, sfgo.PROCESS_SRC, Mapper.MapStr(SF_PROC_EXE)(r))
	if err != nil {
		return err
	}
	r.Ctx.AddHash(hs)
	return nil
}

func HashMd5FileFunc(r *Record) error {
	hs, err := computeHash(crypto.MD5, sfgo.FILE_SRC, Mapper.MapStr(SF_FILE_PATH)(r))
	if err != nil {
		return err
	}
	r.Ctx.AddHash(hs)
	return nil
}

func HashSha1ProcFunc(r *Record) error {
	hs, err := computeHash(crypto.SHA1, sfgo.PROCESS_SRC, Mapper.MapStr(SF_PROC_EXE)(r))
	if err != nil {
		return err
	}
	r.Ctx.AddHash(hs)
	return nil
}

func HashSha1FileFunc(r *Record) error {
	hs, err := computeHash(crypto.SHA1, sfgo.FILE_SRC, Mapper.MapStr(SF_FILE_PATH)(r))
	if err != nil {
		return err
	}
	r.Ctx.AddHash(hs)
	return nil
}

func HashSha256ProcFunc(r *Record) error {
	hs, err := computeHash(crypto.SHA256, sfgo.PROCESS_SRC, Mapper.MapStr(SF_PROC_EXE)(r))
	if err != nil {
		return err
	}
	r.Ctx.AddHash(hs)
	return nil
}

func HashSha256FileFunc(r *Record) error {
	hs, err := computeHash(crypto.SHA256, sfgo.FILE_SRC, Mapper.MapStr(SF_FILE_PATH)(r))
	if err != nil {
		return err
	}
	r.Ctx.AddHash(hs)
	return nil
}

// Utility function for all hash actions
func computeHash(hash crypto.Hash, source sfgo.Source, filePath string) (*HashSet, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	h := hash.New()
	if _, err := io.Copy(h, file); err != nil {
		return nil, err
	}

	hs := new(HashSet)
	hs.Source = source
	hs.Algorithm = hash
	hs.Value = hex.EncodeToString(h.Sum(nil))
	return hs, nil
}
