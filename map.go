package kmap

type Map[K comparable, V comparable] interface {
	Get(key K) V
	GetOrDefault(key K, defaultValue V) V
	Put(key K, value V)
	Clear()
	ContainsKey(key K) bool
	ContainsValue(value V) bool
	Keys() []K
	Values() []V
	Remove(key K)
	Size() int
	IsEmpty() bool
}
