package maps

import (
	"sync"
)

type Slice struct {
	sync.RWMutex
	data map[interface{}][]interface{}
}

func NewSlice() *Slice {
	return &Slice{
		data: make(map[interface{}][]interface{}),
	}
}

func (s *Slice) Store(key interface{}, value ...interface{}) {
	s.Lock()
	defer s.Unlock()
	slice, ok := s.data[key]
	if !ok || len(slice) == 0 {
		s.data[key] = value
		return
	}
	newSlice := make([]interface{}, len(slice)+len(value))
	copy(newSlice, slice)
	copy(newSlice[len(slice):], value)
	delete(s.data, key)
	s.data[key] = newSlice
}

func (s *Slice) Load(key interface{}) (value []interface{}, ok bool) {
	s.RLock()
	defer s.RUnlock()
	value, ok = s.data[key]
	return
}

func (s *Slice) Delete(key interface{}) {
	s.Lock()
	defer s.Unlock()
	delete(s.data, key)
}

func (s *Slice) Range(f func(key interface{}, value []interface{}) bool) {
	s.RLock()
	defer s.RUnlock()
	for k, v := range s.data {
		if !f(k, v) {
			break
		}
	}
}

func (s *Slice) Len() int {
	return len(s.data)
}

func (s *Slice) Clear() {
	s.Lock()
	defer s.Unlock()
	clear(s.data)
}
