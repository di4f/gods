package poolx

// The package implements ordered
// pool structure without any indexex.
// Should be used with only-one-value based
// structures.

type Pool[V comparable] struct {
	store map[V] uint64
	last uint64
}

// Returns new empty pool.
func New[V comparable]() *Pool {
	return &Pool{
		make(map[V] uint64),
		0,
	}
}

func (p *Pool[V]) Push(v V) {
	p.last++
	map[V] = p.last
}

