package cache

type Slice struct {
	items []item
}

func NewSlice(capHint int) *Slice {
	return &Slice{
		items: make([]item, 0, capHint),
	}
}

func (s *Slice) Get(id string) interface{} {
	for _, it := range s.items {
		if it.id == id {
			return it.value
		}
	}
	return nil
}

func (s *Slice) Set(id string, value interface{}) {
	for i, it := range s.items {
		if it.id == id {
			s.items[i].value = value
			return
		}
	}
	s.items = append(s.items, item{id, value})
}

func (s *Slice) Len() int {
	return len(s.items)
}

var _ Cache = (*Slice)(nil)
