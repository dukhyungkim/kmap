package kmap

type sortedMap[K comparable, V comparable] struct {
	keys  []K
	store map[K]V
}

func NewSortedMap[K comparable, V comparable]() Map[K, V] {
	return &sortedMap[K, V]{
		store: make(map[K]V),
	}
}

func (m *sortedMap[K, V]) Get(key K) V {
	return m.store[key]
}

func (m *sortedMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	v, has := m.store[key]
	if !has {
		return defaultValue
	}
	return v
}

func (m *sortedMap[K, V]) Put(key K, value V) {
	m.store[key] = value
	m.keys = append(m.keys, key)
}

func (m *sortedMap[K, V]) Clear() {
	m.store = make(map[K]V)
	m.keys = []K{}
}

func (m *sortedMap[K, V]) ContainsKey(key K) bool {
	_, has := m.store[key]
	return has
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

func (m *sortedMap[K, V]) Values() []V {
	values := make([]V, 0, len(m.store))
	for _, v := range m.store {
		values = append(values, v)
	}
	return values
}

func (m *sortedMap[K, V]) Remove(key K) {
	delete(m.store, key)
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

func (m *sortedMap[K, V]) IsEmpty() bool {
	return m.Size() == 0
}
