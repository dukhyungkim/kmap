package kmap

type linkedHashMap[K KeyAble, V comparable] struct {
	keys  []K
	store map[K]V
}

func NewLinkedHashMap[K KeyAble, V comparable]() Map[K, V] {
	return &linkedHashMap[K, V]{
		store: make(map[K]V),
	}
}

func (m *linkedHashMap[K, V]) Get(key K) V {
	return m.store[key]
}

func (m *linkedHashMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	v, has := m.store[key]
	if !has {
		return defaultValue
	}
	return v
}

func (m *linkedHashMap[K, V]) Put(key K, value V) {
	if _, has := m.store[key]; !has {
		m.keys = append(m.keys, key)
	}
	m.store[key] = value
}

func (m *linkedHashMap[K, V]) Clear() {
	m.store = make(map[K]V)
	m.keys = []K{}
}

func (m *linkedHashMap[K, V]) ContainsKey(key K) bool {
	_, has := m.store[key]
	return has
}

func (m *linkedHashMap[K, V]) ContainsValue(value V) bool {
	for _, v := range m.Values() {
		if v == value {
			return true
		}
	}
	return false
}

func (m *linkedHashMap[K, V]) Keys() []K {
	return m.keys
}

func (m *linkedHashMap[K, V]) Values() []V {
	values := make([]V, 0, len(m.store))
	for _, v := range m.store {
		values = append(values, v)
	}
	return values
}

func (m *linkedHashMap[K, V]) Remove(key K) {
	delete(m.store, key)
	for i, k := range m.keys {
		if k == key {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
			break
		}
	}
}

func (m *linkedHashMap[K, V]) Size() int {
	return len(m.keys)
}

func (m *linkedHashMap[K, V]) IsEmpty() bool {
	return m.Size() == 0
}
