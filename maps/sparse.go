package maps

import (
	"slices"
	"cmp"
	"fmt"
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
type mSparse[K cmp.Ordered, V any] struct {
	store map[K] V
	def V
	keys []K
}

// Return the new with the default implementation storing the values from the
// map inside the structure. Is fast on creation and reading, but
// slow on inserting and deleting. Takes only one or zero maps as input.
// Panics otherwise.
func NewSparse[K cmp.Ordered, V any](
	defval V,
	values ...map[K] V,
) Map[K, V] {

	var (
		store map[K] V
		keys []K
		vs map[K] V
	)
	if len(values) > 1 {
		panic("too many arguments")
	} else if len(values) == 1 {
		vs = values[0]
	}

	if vs == nil {
		store = map[K] V{}
		keys = []K{}
	} else {
		store = make(map[K] V, len(vs))
		keys = make([]K, len(vs))
		i := 0
		for k, v := range vs {
			keys[i] = k
			store[k] = v
			i++
		}
		slices.Sort(keys)
	}

	ret := &mSparse[K, V]{
		store: store,
		def: defval,
		keys: keys,
	}

	return ret
}

func (s *mSparse[K, V]) Size() int {
	return len(s.keys)
}

func (s *mSparse[K, V]) Clear() {
	s.store = map[K]V{}
	s.keys = []K{}
}

func (s *mSparse[K, V]) Has(key K) bool {
	_, ok := s.store[key]
	return ok
}

func (s *mSparse[K, V]) Get(key K) (V) {
	val, ok := s.store[key]
	if !ok {
		val = s.def
	}
	return val
}

func (s *mSparse[K, V]) Got(key K) (V, bool) {
	v, ok := s.store[key]
	if !ok {
		v = s.def
	}
	return v, ok
}

func (s *mSparse[K, V]) Set(k K, v V) {
	_, ok := s.store[k]
	if !ok {
		s.keys = append(s.keys, k)
		s.sort()
	}

	s.store[k] = v
}

func (s *mSparse[K, V]) Empty() bool {
	return len(s.keys) == 0
}

func (s *mSparse[K, V]) Del(k K) {
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
		// Removing the key.
		s.keys = append(s.keys[:idx], s.keys[idx+1:]...)
	}
}

func (s *mSparse[K, V]) Chan(
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

func (s *mSparse[K, V]) KeyChan() chan K {
	ret := make(chan K)
	go func() {
		for _, k := range s.keys {
			ret <- k
		}
		close(ret)
	}()
	return ret
}

func (s *mSparse[K, V]) Keys() []K {
	return s.keys
}

func (s *mSparse[K, V]) Values() []V {
	ret := []V{}
	for v := range s.Chan() {
		ret = append(ret, v)
	}
	return ret
}

func (s *mSparse[K, V]) String() string {
	return fmt.Sprintf("%#v", s.store)
}

func (s *mSparse[K, V]) sort() {
	slices.Sort(s.keys)
}

