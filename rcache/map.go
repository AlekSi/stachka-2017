package rcache

import (
	"sync"
)

type Map struct {
	m     sync.RWMutex
	items map[string]interface{}
}

func NewMap(capHint int) *Map {
	return &Map{
		items: make(map[string]interface{}, capHint),
	}
}

func (m *Map) Get(id string) interface{} {
	m.m.RLock()
	v := m.items[id]
	m.m.RUnlock()
	return v
}

func (m *Map) Set(id string, value interface{}) {
	m.m.Lock()
	m.items[id] = value
	m.m.Unlock()
}

func (m *Map) Len() int {
	m.m.Lock()
	l := len(m.items)
	m.m.Unlock()
	return l
}

var _ Cache = (*Map)(nil)
