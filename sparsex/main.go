package sparsex

import (
	"sort"
	cons "golang.org/x/exp/constraints"
)

// The package implements sparse array structure.
// You can make sparse matrices by assigning
// the value type to another sparse array.

// The sparse array type.
// Does not sort by default on setting so by
// default is not ordered for the Chan, Keys, KeyChan, Slice methods.
// Its made for the optimization sakes because
// initially the package was made for the gox (2D game engine) package use
// to provide dynamic layering.
// See *Sparse[K, V].ShouldSort method to change it.
type Sparse[K cons.Ordered, V any] struct {
	store map[K] V
	def V
	keys []K
	shouldSort bool
}

// Returns new sparse array taking all the values from valuesStores
// into the structure.
func New[K cons.Ordered, V any](valueStores ...map[K] V) *Sparse[K, V] {

	finalLen := 0
	for _, valueStore := range valueStores {
		finalLen = len(valueStore)
	}
	keys := make([]K, finalLen)
	store := make(map[K] V)
	for _, valueStore := range valueStores {
		i := 0
		for k, v := range valueStore {
			keys[i] = k
			store[k] = v
			i++
		}
	}
	return &Sparse[K, V]{
		store: store,
		keys: keys,
	}
}

// Define if should sort each time we set or delete.
func (s *Sparse[K, V]) ShouldSort(v bool) *Sparse[K, V] {
	s.shouldSort = v
	if s.shouldSort {
		s.Sort()
	}
	return s
}

// Sets the default sparse array value.
// Returned when there is no key for the value.
func (s *Sparse[K, V]) Default(v V) *Sparse[K, V] {
	s.def = v
	return s
}

// Get the value by the key. The secound value
// represents whether the array contains the value.
// If the array does not contain the value then default one
// will be returned.
func (s *Sparse[K, V]) Get(key K) (V, bool) {
	val, ok := s.store[key]
	if !ok {
		val = s.def
	}
	return val, ok
}

// Set the value to the key.
func (s *Sparse[K, V]) Set(k K, v V) {
	_, ok := s.store[k]
	if !ok {
		s.keys = append(s.keys, k)
		if s.shouldSort {
			s.Sort()
		}
	}

	s.store[k] = v
}

// Delete the value by the key.
func (s *Sparse[K, V]) Delete(k K) {
	delete(s.store, k)

	// To know if the loop was run.
	idx := -1

	for i, v := range s.keys {
		if v == k {
			idx = i
			break
		}
	}

	if idx != -1 {
		s.keys = append(s.keys[:idx], s.keys[idx+1:]...)
		if s.shouldSort {
			s.Sort()
		}
	}
}

// Alias for the Delete method.
func (s *Sparse[K, V]) Del(k K) {
	s.Delete(k)
}

// Returns channel of pairs.
func (s *Sparse[K, V]) Chan(
) chan V {
	keys := s.keys
	store := s.store
	ret := make(chan V)

	go func() {
		for _, k := range keys {
			ret <- store[k]
		}
		close(ret)
	}()

	return ret
}

// Returns a slice of the keys of the array.
func (s *Sparse[K, V]) Keys() []K {
	return s.keys
}

// Returns channel of keys in its order in the structure.
func (s *Sparse[K, V]) KeyChan() chan K {
	ret := make(chan K)
	go func() {
		for _, k := range s.keys {
			ret <- k
		}
		close(ret)
	}()
	return ret
}

// Returns slice of the already set values.
func (s *Sparse[K, V]) Slice() []V {
	ret := []V{}
	for v := range s.Chan() {
		ret = append(ret, v)
	}
	return ret
}

// Sort the keys making current array finely ordered.
func (s *Sparse[K, V]) Sort() {
	sort.Slice(s.keys, func(i, j int) bool {
		return s.keys[i] < s.keys[j]
	})
}

