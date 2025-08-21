package maps

import (
	"sync"
)

type Map struct {
	sync.RWMutex
	data map[interface{}]map[interface{}]interface{}
}

func NewMap() *Map {
	return &Map{
		data: make(map[interface{}]map[interface{}]interface{}),
	}
}

func (m *Map) Store(key interface{}, value map[interface{}]interface{}) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.data[key]; !ok {
		m.data[key] = make(map[interface{}]interface{})
		m.data[key] = value
		return
	}
	for k, v := range value {
		m.data[key][k] = v
	}
}

func (m *Map) Load(key interface{}) (value map[interface{}]interface{}, ok bool) {
	m.RLock()
	defer m.RUnlock()
	value, ok = m.data[key]
	return
}

func (m *Map) Delete(key interface{}) {
	m.Lock()
	defer m.Unlock()
	delete(m.data, key)
}

func (m *Map) Range(f func(key interface{}, value map[interface{}]interface{}) bool) {
	m.RLock()
	defer m.RUnlock()
	for k, v := range m.data {
		if !f(k, v) {
			break
		}
	}
}

func (m *Map) Len() int {
	return len(m.data)
}

func (m *Map) Clear() {
	m.Lock()
	defer m.Unlock()
	clear(m.data)
}
