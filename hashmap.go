package kmap

type hashMap[K comparable, V any] struct {
	store map[K]V
}

func NewHashMap[K comparable, V any]() Map[K, V] {
	return &hashMap[K, V]{
		store: make(map[K]V),
	}
}

func (m *hashMap[K, V]) Get(key K) V {
	return m.store[key]
}

func (m *hashMap[K, V]) Put(key K, value V) {
	m.store[key] = value
}
