package poolx

import (
	"github.com/surdeus/godat/src/llx"
)

// Ordered value-only based structure.
// Fast deleting by value.
// Cannot store multiple the same values.

type Pool[V comparable] struct {
	store *llx.LinkedList[V]
	keys map[V] int
}

// Returns new empty pool.
func New[V comparable]() *Pool {
	return &Pool{
		llx.New[V]()
		0,
	}
}

func (p *Pool[V]) Append(v V) {
	p.store.Append(v)
	
}

