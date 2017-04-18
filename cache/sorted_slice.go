package cache

import (
	"sort"
)

type SortedSlice struct {
	items []item
}

func NewSortedSlice(capHint int) *SortedSlice {
	return &SortedSlice{
		items: make([]item, 0, capHint),
	}
}

func (s *SortedSlice) Get(id string) interface{} {
	i := sort.Search(len(s.items), func(i int) bool { return s.items[i].id >= id })
	if i < len(s.items) && s.items[i].id == id {
		return s.items[i].value
	}
	return nil
}

func (s *SortedSlice) Set(id string, value interface{}) {
	i := sort.Search(len(s.items), func(i int) bool { return s.items[i].id >= id })
	if i == len(s.items) {
		s.items = append(s.items, item{id, value})
		return
	}
	if s.items[i].id == id {
		s.items[i].value = value
		return
	}

	s.items = append(s.items, item{})
	copy(s.items[i+1:], s.items[i:])
	s.items[i] = item{id, value}
}

func (s *SortedSlice) Len() int {
	return len(s.items)
}

var _ Cache = (*SortedSlice)(nil)
