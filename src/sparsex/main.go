package sparsex

import (
	"sort"
	cons "golang.org/x/exp/constraints"
)

// The package implements a simple ordered map.
// In fact can be used as a sparse array so it is
// where the name comes from.

type Pair[K cons.Ordered, V any] struct {
	K K
	V V
}

type Sparse[K cons.Ordered, V any] struct {
	store map[K] V
	keys []K
	shouldSort bool
}

// Returns new sparse array
func New[K cons.Ordered, V any](s bool) *Sparse[K, V] {
	return &Sparse[K, V]{
		store: make(map[K] V),
		keys: []K{},
		shouldSort: s,
	}
}


func (s *Sparse[K, V]) Get(key K) (V, bool) {
	val, ok := s.store[key]
	return val, ok
}

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

func (s Sparse[K, V]) Del(k K) {
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

func (s *Sparse[K, V]) Vals(
) chan Pair[K, V] {
	keys := s.keys
	store := s.store
	ret := make(chan Pair[K, V])

	go func() {
		for _, v := range keys {
			ret <- Pair[K, V]{
				v, store[v],
			}
		}
		close(ret)
	}()

	return ret
}

// Sort the keys.
func (s *Sparse[K, V]) Sort() {
	sort.Slice(s.keys, func(i, j int) bool {
		return s.keys[i] < s.keys[j]
	})
}

