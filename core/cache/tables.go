//
// Copyright (C) 2020 IBM Corporation.
//
// Authors:
// Frederico Araujo <frederico.araujo@ibm.com>
// Teryl Taylor <terylt@ibm.com>
//
package cache

import (
	"crypto/sha256"
	"fmt"
	"sync"

	cqueue "github.com/enriquebris/goconcurrentqueue"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/sysflow-telemetry/sf-apis/go/sfgo"
	"github.ibm.com/sysflow/goutils/logger"
)

const (
	cacheSize = 2
)

var instance *SFTables
var once sync.Once

// SFTables defines thread-safe shared cache for plugins for storing SysFlow entities.
type SFTables struct {
	contTable *cqueue.FIFO
	procTable *cqueue.FIFO
	fileTable *cqueue.FIFO
	rwmutex   sync.RWMutex
	capacity  int
}

// GetInstance returns SFTables singleton instance
func GetInstance() *SFTables {
	once.Do(func() {
		instance = newSFTables(cacheSize)
	})
	return instance
}

// newSFTables creates a new SFTables instance.
func newSFTables(capacity int) *SFTables {
	t := new(SFTables)
	if capacity < 1 {
		logger.Error.Println("Cache capacity must be greater than 1")
		return nil
	}
	t.capacity = capacity
	t.contTable = cqueue.NewFIFO()
	t.procTable = cqueue.NewFIFO()
	t.fileTable = cqueue.NewFIFO()
	t.contTable.Enqueue(cmap.New())
	t.procTable.Enqueue(cmap.New())
	t.fileTable.Enqueue(cmap.New())
	return t
}

// Reset pushes a new set of empty maps into the cache.
func (t *SFTables) Reset() {
	t.rwmutex.Lock()
	defer t.rwmutex.Unlock()
	t.reset(t.contTable)
	t.reset(t.procTable)
	t.reset(t.fileTable)
}

func (t *SFTables) reset(queue *cqueue.FIFO) {
	queue.Enqueue(cmap.New())
	if queue.GetLen() > t.capacity {
		queue.Remove(0)
	}
}

// GetCont retrieves a cached container object by ID.
func (t *SFTables) GetCont(ID string) *sfgo.Container {
	t.rwmutex.RLock()
	defer t.rwmutex.RUnlock()
	for i := 0; i < t.contTable.GetLen(); i++ {
		m, _ := t.contTable.Get(i)
		table := m.(cmap.ConcurrentMap)
		if v, ok := table.Get(ID); ok {
			return v.(*sfgo.Container)
		}
	}
	return nil
}

// SetCont stores a container object in the cache.
func (t *SFTables) SetCont(ID string, o *sfgo.Container) {
	t.rwmutex.RLock()
	m, _ := t.contTable.Get(t.contTable.GetLen() - 1)
	t.rwmutex.RUnlock()
	table := m.(cmap.ConcurrentMap)
	table.Set(ID, o)
}

// GetProc retrieves a cached process object by ID.
func (t *SFTables) GetProc(ID sfgo.OID) *sfgo.Process {
	t.rwmutex.RLock()
	defer t.rwmutex.RUnlock()
	for i := 0; i < t.procTable.GetLen(); i++ {
		m, _ := t.procTable.Get(i)
		table := m.(cmap.ConcurrentMap)
		if v, ok := table.Get(t.getHash(ID)); ok {
			return v.(*sfgo.Process)
		}
	}
	return nil
}

// SetProc stores a process object in the cache.
func (t *SFTables) SetProc(ID sfgo.OID, o *sfgo.Process) {
	t.rwmutex.RLock()
	m, _ := t.procTable.Get(t.procTable.GetLen() - 1)
	t.rwmutex.RUnlock()
	table := m.(cmap.ConcurrentMap)
	table.Set(t.getHash(ID), o)
}

// GetFile retrieves a cached file object by ID.
func (t *SFTables) GetFile(ID sfgo.FOID) *sfgo.File {
	t.rwmutex.RLock()
	defer t.rwmutex.RUnlock()
	for i := 0; i < t.fileTable.GetLen(); i++ {
		m, _ := t.fileTable.Get(i)
		table := m.(cmap.ConcurrentMap)
		if v, ok := table.Get(t.getHash(ID)); ok {
			return v.(*sfgo.File)
		}
	}
	return nil
}

// SetFile stores a file object in the cache.
func (t *SFTables) SetFile(ID sfgo.FOID, o *sfgo.File) {
	t.rwmutex.RLock()
	m, _ := t.fileTable.Get(t.fileTable.GetLen() - 1)
	t.rwmutex.RUnlock()
	table := m.(cmap.ConcurrentMap)
	table.Set(t.getHash(ID), o)
}

func (t *SFTables) getHash(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))
	return fmt.Sprintf("%x", h.Sum(nil))
}
