package cache

type Map struct {
	items map[string]interface{}
}

func NewMap(capHint int) *Map {
	return &Map{
		items: make(map[string]interface{}, capHint),
	}
}

func (m *Map) Get(id string) interface{} {
	return m.items[id]
}

func (m *Map) Set(id string, value interface{}) {
	m.items[id] = value
}

func (m *Map) Len() int {
	return len(m.items)
}

var _ Cache = (*Map)(nil)
