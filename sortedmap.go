package kmap

import "sort"

type sortedMap[K KeyAble, V comparable] struct {
	keys []K
	Map[K, V]
}

func NewSortedMap[K KeyAble, V comparable]() Map[K, V] {
	return &sortedMap[K, V]{
		Map: NewHashMap[K, V](),
	}
}

func (m *sortedMap[K, V]) Put(key K, value V) {
	if !m.Map.ContainsKey(key) {
		m.keys = append(m.keys, key)
		sort.Slice(m.keys, func(i, j int) bool {
			return m.keys[i] < m.keys[j]
		})
	}
	m.Map.Put(key, value)
}

func (m *sortedMap[K, V]) Clear() {
	m.Map.Clear()
	m.keys = []K{}
}

func (m *sortedMap[K, V]) ContainsValue(value V) bool {
	for _, v := range m.Values() {
		if v == value {
			return true
		}
	}
	return false
}

func (m *sortedMap[K, V]) Keys() []K {
	return m.keys
}

func (m *sortedMap[K, V]) Remove(key K) {
	m.Map.Remove(key)
	for i, k := range m.keys {
		if k == key {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
			break
		}
	}
}

func (m *sortedMap[K, V]) Size() int {
	return len(m.keys)
}
