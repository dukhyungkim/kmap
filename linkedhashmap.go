package kmap

type linkedHashMap[K KeyAble, V comparable] struct {
	keys []K
	Map[K, V]
}

func NewLinkedHashMap[K KeyAble, V comparable]() Map[K, V] {
	return &linkedHashMap[K, V]{
		Map: NewHashMap[K, V](),
	}
}

func (m *linkedHashMap[K, V]) Put(key K, value V) {
	if !m.Map.ContainsKey(key) {
		m.keys = append(m.keys, key)
	}
	m.Map.Put(key, value)
}

func (m *linkedHashMap[K, V]) Clear() {
	m.Map.Clear()
	m.keys = []K{}
}

func (m *linkedHashMap[K, V]) Keys() []K {
	return m.keys
}

func (m *linkedHashMap[K, V]) Remove(key K) {
	m.Map.Remove(key)
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
