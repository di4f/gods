package poolx

import (
	"github.com/surdeus/godat/src/llx"
)

// Ordered value-only based structure.
// Fast deleting by value.
// Cannot store multiple equal values.

type Pool[V comparable] struct {
	store *llx.LinkedList[V]
}

// Return new empty pool.
func New[V comparable]() *Pool[V] {
	return &Pool[V]{
		store: llx.New[V](),
	}
}

func (p *Pool[V]) Append(v V) {
	p.store.Append(v)
}

// Deletes the first appearance of the value in the list.
func (p *Pool[V]) Del(v V) bool {
	i := 0
	ll := p.store
	for e := ll.First() ; e != nil ; e = e.Next() {
		if e.Value() == v {
			ll.Del(i)
			return true
		}
		
		i++
	}
	
	return false
}

func (p *Pool[V]) Range() chan llx.Pair[V] {
	return p.store.Range()
}

