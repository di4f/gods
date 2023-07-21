package sparsex

import (
	"sort"
	cons "golang.org/x/exp/constraints"
	"github.com/mojosa-software/godat/src/iterx"
)

// The package implements a simple ordered map.
// In fact can be used as a sparse array so it is
// where the name comes from.

// The sparse array type.
type Sparse[K cons.Ordered, V any] struct {
	store map[K] V
	keys []K
	shouldSort bool
}

// Returns new sparse array.
// If shouldSort == true then it will sort the array on
// each change.
func New[K cons.Ordered, V any](shouldSort bool) *Sparse[K, V] {
	return &Sparse[K, V]{
		store: make(map[K] V),
		keys: []K{},
		shouldSort: shouldSort,
	}
}

// Get the value by the key.
func (s *Sparse[K, V]) Get(key K) (V, bool) {
	val, ok := s.store[key]
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

// Returns channel of pairs.
func (s *Sparse[K, V]) Chan(
) iterx.PairChan[K, V] {
	keys := s.keys
	store := s.store
	ret := make(iterx.PairChan[K, V])

	go func() {
		for _, k := range keys {
			ret <- iterx.Pair[K, V]{
				K: k,
				V: store[k],
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

