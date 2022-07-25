package kmap

type hashMap[K comparable, V comparable] struct {
	store map[K]V
}

func NewHashMap[K comparable, V comparable]() Map[K, V] {
	return &hashMap[K, V]{
		store: make(map[K]V),
	}
}

func (m *hashMap[K, V]) Get(key K) V {
	return m.store[key]
}

func (m *hashMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	v, has := m.store[key]
	if !has {
		return defaultValue
	}
	return v
}

func (m *hashMap[K, V]) Put(key K, value V) {
	m.store[key] = value
}

func (m *hashMap[K, V]) Clear() {
	m.store = make(map[K]V)
}

func (m *hashMap[K, V]) ContainsKey(key K) bool {
	_, has := m.store[key]
	return has
}

func (m *hashMap[K, V]) ContainsValue(value V) bool {
	for _, v := range m.Values() {
		if v == value {
			return true
		}
	}
	return false
}

func (m *hashMap[K, V]) Keys() []K {
	keys := make([]K, 0, len(m.store))
	for k := range m.store {
		keys = append(keys, k)
	}
	return keys
}

func (m *hashMap[K, V]) Values() []V {
	values := make([]V, 0, len(m.store))
	for _, v := range m.store {
		values = append(values, v)
	}
	return values
}

func (m *hashMap[K, V]) Remove(key K) {
	delete(m.store, key)
}

func (m *hashMap[K, V]) Size() int {
	return len(m.store)
}

func (m *hashMap[K, V]) IsEmpty() bool {
	return m.Size() == 0
}
