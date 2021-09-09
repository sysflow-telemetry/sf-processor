//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
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

// Package cache implements a singleton cache for telemetry objects.
package cache

import (
	"sync"

	"github.com/sysflow-telemetry/sf-apis/go/hash"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)

var instance *SFTables
var once sync.Once

// SFTables defines thread-safe shared cache for plugins for storing SysFlow entities.
type SFTables struct {
	contTable map[string]*sfgo.Container
	procTable map[uint64][]*sfgo.Process
	fileTable map[uint64]*sfgo.File
	rwmutex   sync.RWMutex
}

// GetInstance returns SFTables singleton instance
func GetInstance() *SFTables {
	once.Do(func() {
		instance = newSFTables()
	})
	return instance
}

// newSFTables creates a new SFTables instance.
func newSFTables() *SFTables {
	t := new(SFTables)
	t.new()
	return t
}

// Reset pushes a new set of empty maps into the cache.
func (t *SFTables) Reset() {
	t.rwmutex.Lock()
	defer t.rwmutex.Unlock()
	t.new()
}

func (t *SFTables) new() {
	t.contTable = make(map[string]*sfgo.Container)
	t.procTable = make(map[uint64][]*sfgo.Process)
	t.fileTable = make(map[uint64]*sfgo.File)
}

// GetCont retrieves a cached container object by ID.
func (t *SFTables) GetCont(ID string) *sfgo.Container {
	t.rwmutex.RLock()
	defer t.rwmutex.RUnlock()
	if v, ok := t.contTable[ID]; ok {
		return v
	}
	return nil
}

// SetCont stores a container object in the cache.
func (t *SFTables) SetCont(ID string, o *sfgo.Container) {
	t.rwmutex.Lock()
	defer t.rwmutex.Unlock()
	t.contTable[ID] = o
}

// GetProc retrieves a cached process object by ID.
func (t *SFTables) GetProc(ID sfgo.OID) *sfgo.Process {
	t.rwmutex.RLock()
	defer t.rwmutex.RUnlock()
	if p, ok := t.procTable[hash.GetHash(ID)]; ok {
		if po := p[sfgo.SFObjectStateMODIFIED]; po != nil {
			return po
		}
		if po := p[sfgo.SFObjectStateCREATED]; po != nil {
			return po
		}
		if po := p[sfgo.SFObjectStateREUP]; po != nil {
			return po
		}
	}
	return nil
}

// SetProc stores a process object in the cache.
func (t *SFTables) SetProc(ID sfgo.OID, o *sfgo.Process) {
	t.rwmutex.Lock()
	defer t.rwmutex.Unlock()
	oid := hash.GetHash(ID)
	if p, ok := t.procTable[oid]; ok {
		p[o.State] = o
	} else {
		p = make([]*sfgo.Process, sfgo.SFObjectStateREUP+1)
		p[o.State] = o
		t.procTable[oid] = p
	}
}

// GetFile retrieves a cached file object by ID.
func (t *SFTables) GetFile(ID sfgo.FOID) *sfgo.File {
	t.rwmutex.RLock()
	defer t.rwmutex.RUnlock()
	if v, ok := t.fileTable[hash.GetHash(ID)]; ok {
		return v
	}
	return nil
}

// SetFile stores a file object in the cache.
func (t *SFTables) SetFile(ID sfgo.FOID, o *sfgo.File) {
	t.rwmutex.Lock()
	defer t.rwmutex.Unlock()
	t.fileTable[hash.GetHash(ID)] = o
}
