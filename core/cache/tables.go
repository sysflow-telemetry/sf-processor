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

// Package cache implements a local cache for telemetry objects.
package cache

import (
	"github.com/sysflow-telemetry/sf-apis/go/hash"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
)

// SFTables defines thread-safe shared cache for plugins for storing SysFlow entities.
type SFTables struct {
	contTable  map[string]*sfgo.Container
	procTable  map[uint64][]*sfgo.Process
	fileTable  map[uint64]*sfgo.File
	ptreeTable map[uint64][]*sfgo.Process
}

// NewSFTables creates a new SFTables instance.
func NewSFTables() *SFTables {
	t := new(SFTables)
	t.new()
	return t
}

func (t *SFTables) new() {
	t.contTable = make(map[string]*sfgo.Container)
	t.procTable = make(map[uint64][]*sfgo.Process)
	t.fileTable = make(map[uint64]*sfgo.File)
	t.ptreeTable = make(map[uint64][]*sfgo.Process)
}

// Reset pushes a new set of empty maps into the cache.
func (t *SFTables) Reset() {
	t.new()
}

// GetCont retrieves a cached container object by ID.
func (t *SFTables) GetCont(ID string) (co *sfgo.Container) {
	co = t.contTable[ID]
	return
}

// SetCont stores a container object in the cache.
func (t *SFTables) SetCont(ID string, o *sfgo.Container) {
	t.contTable[ID] = o
}

// GetProc retrieves a cached process object by ID.
func (t *SFTables) GetProc(ID sfgo.OID) (po *sfgo.Process) {
	if p, ok := t.procTable[hash.GetHash(ID)]; ok {
		if v := p[sfgo.SFObjectStateMODIFIED]; v != nil {
			po = v
		} else if v := p[sfgo.SFObjectStateCREATED]; v != nil {
			po = v
		} else if v := p[sfgo.SFObjectStateREUP]; v != nil {
			po = v
		}
	}
	return
}

// SetProc stores a process object in the cache.
func (t *SFTables) SetProc(ID sfgo.OID, o *sfgo.Process) {
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
	if v, ok := t.fileTable[hash.GetHash(ID)]; ok {
		return v
	}
	return nil
}

// SetFile stores a file object in the cache.
func (t *SFTables) SetFile(ID sfgo.FOID, o *sfgo.File) {
	t.fileTable[hash.GetHash(ID)] = o
}

// GetPtree retrieves and caches the processes hierachy given a process ID.
func (t *SFTables) GetPtree(ID sfgo.OID) []*sfgo.Process {
	oID := hash.GetHash(ID)
	if ptree, ok := t.ptreeTable[oID]; ok {
		return ptree
	}
	ptree := t.getProcProv(ID)
	t.ptreeTable[oID] = ptree
	return ptree
}

// getProcProv builds the provenance tree of a process recursevely.
func (t *SFTables) getProcProv(ID sfgo.OID) []*sfgo.Process {
	var ptree = make([]*sfgo.Process, 0)
	if p := t.GetProc(ID); p != nil {
		if p.Poid != nil && p.Poid.UnionType == sfgo.UnionNullOIDTypeEnumOID {
			return append(append(ptree, p), t.getProcProv(*p.Poid.OID)...)
		}
		return append(ptree, p)
	}
	return ptree
}
