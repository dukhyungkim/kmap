package kmap

type sortedMap[K comparable, V comparable] struct {
}

func NewSortedMap[K comparable, V comparable]() Map[K, V] {
	return &sortedMap[K, V]{}
}

func (s *sortedMap[K, V]) Get(key K) V {
	//TODO implement me
	panic("implement me")
}

func (s *sortedMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	//TODO implement me
	panic("implement me")
}

func (s *sortedMap[K, V]) Put(key K, value V) {
	//TODO implement me
	panic("implement me")
}

func (s *sortedMap[K, V]) PutIfAbsent(key K, value V) V {
	//TODO implement me
	panic("implement me")
}

func (s *sortedMap[K, V]) Clear() {
	//TODO implement me
	panic("implement me")
}

func (s *sortedMap[K, V]) ContainsKey(key K) bool {
	//TODO implement me
	panic("implement me")
}

func (s *sortedMap[K, V]) ContainsValue(value V) bool {
	//TODO implement me
	panic("implement me")
}

func (s *sortedMap[K, V]) Keys() []K {
	//TODO implement me
	panic("implement me")
}

func (s *sortedMap[K, V]) Values() []V {
	//TODO implement me
	panic("implement me")
}

func (s *sortedMap[K, V]) Remove(key K) {
	//TODO implement me
	panic("implement me")
}

func (s *sortedMap[K, V]) Size() int {
	//TODO implement me
	panic("implement me")
}

func (s *sortedMap[K, V]) IsEmpty() bool {
	//TODO implement me
	panic("implement me")
}
