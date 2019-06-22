package cache

import (
	"sync"
)

type SyncMap struct {
	items sync.Map
}

func NewSyncMap() *SyncMap {
	return new(SyncMap)
}

func (m *SyncMap) Get(id string) interface{} {
	v, _ := m.items.Load(id)
	return v
}

func (m *SyncMap) Set(id string, value interface{}) {
	m.items.Store(id, value)
}

func (m *SyncMap) Len() int {
	var c int
	m.items.Range(func(key, value interface{}) bool {
		c++
		return true
	})
	return c
}

var _ Cache = (*SyncMap)(nil)
