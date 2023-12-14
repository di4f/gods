package maps

import (
	"fmt"
	"github.com/di4f/gods"
)

// Generic map interface for all the maps.
type Map[K comparable, V any] interface {
	gods.Container[V]

	// Returns if the map has the key set.
	Has(K) bool

	// Get the value and if the key is not set panic.
	Get(K) V

	// The Go way to get values without panicing.
	Got(K) (V, bool)

	// Set the value or reset if it is already set.
	Set(K, V)

	// Delete the key no matter it exists or not.
	Del(K)

	// Returns slice of values.
	// For the order look the comment
	// for "Keys()".
	Values() []V

	// Get the values channel.
	Chan() chan V

	// Returns slice of keys.
	// Order is not guaranteed if
	// the is not specified otherwise
	// like for the NewOrdered.
	Keys() []K

	KeyChan() chan K

	// The function to range over the values
	// with the keys.
	/*Range() chan struct{
		K K
		V V
	}*/
}
type lMap[K comparable, V any] struct {
	store map[K]V
}

// Returns new basic map with the builtin Go type down there.
// Has all the features of the builtin Go maps and same performance.
func New[K comparable, V any]() Map[K, V] {
	ret := &lMap[K, V]{}
	ret.store = map[K]V{}
	return ret
}

func (m *lMap[K, V]) Clear() {
	m.store = map[K]V{}
}

func (m *lMap[K, V]) Keys() []K {
	r := make([]K, len(m.store))
	i := 0
	for k := range m.store {
		r[i] = k
		i++
	}

	return r
}

func (m *lMap[K, V]) Empty() bool {
	return len(m.store) == 0
}

func (m *lMap[K, V]) Del(key K) {
	delete(m.store, key)
}

func (m *lMap[K, V]) Values() []V {
	r := make([]V, len(m.store))
	i := 0
	for _, v := range m.store {
		r[i] = v
		i++
	}

	return r
}

func (m *lMap[K, V]) Chan() chan V {
	ret := make(chan V)
	go func() {
		for _, v := range m.store {
			ret <- v
		}
		close(ret)
	}()
	return ret
}

func (m *lMap[K, V]) Has(k K) bool {
	_, ok := m.store[k]
	return ok
}

func (m *lMap[K, V]) Set(k K, v V) {
	m.store[k] = v
}

func (m *lMap[K, V]) Get(key K) V {
	v, ok := m.store[key]
	if !ok {
		panic(fmt.Sprintf("there is no such key '%v'", key))
	}
	return v
}

func (m *lMap[K, V]) Got(key K) (V, bool) {
	v, ok := m.store[key]
	return v, ok
}

func (m *lMap[K, V]) Size() int {
	return len(m.store)
}

func (m *lMap[K, V]) KeyChan() chan K {
	ret := make(chan K)
	go func() {
		for k := range m.store {
			ret <- k
		}
		close(ret)
	}()
	return ret
}
