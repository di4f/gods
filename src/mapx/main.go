package mapx

import (
	"fmt"
)

// The package implements some more specific
// map structures for special uses.
// Implemented mostly to be embedded to other structures.

// General map type, wrap for the built-in one.
type Map[K comparable, V any] map[K] V

// Returns new basic map.
func New[K comparable, V any]() Map[K, V] {
	return make(Map[K, V])
}

// Returns slice of keys of the map.
func (m Map[K, V]) Keys() []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}

	return r
}

// Returns slice of values of the map.
func (m Map[K, V]) Values(sm map[K] V) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}

	return r
}

// Checks if the map contains the key.
func (m Map[K, V]) Has(k K) bool {
	_, ok := m[k]
	return ok
}

// Sets the new value by key or resets if it exists.
func (m Map[K, V]) Set(k K, v V) {
	m[k] = v
}

// Returns the value by key. Panics if there is no such key.
func (m Map[K, V]) Get(k K) V {
	v, ok := m[k]
	if !ok {
		panic(fmt.Sprintf("there is no such key '%v'", k))
	}
	return v
}

