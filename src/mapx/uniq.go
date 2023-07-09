package mapx

// The type implements map type where
// you can get, set and delete by value
// since it store everything as ONLY entity
// both for keys and values way.
// Use only when you do not care about the order.
type UniqMap[K, V comparable] struct {
	store map[K] V
	rstore map[V] K
}

// Returns new empty UniqMap.
func NewUniq[K, V comparable]() *UniqMap[K, V] {
	return &UniqMap[K, V]{
		make(map[K] V),
		make(map[V] K),
	}
}

// Sets new value v for the k key.
func (m *UniqMap[K, V]) Set(k K, v V) {
	m.store[k] = v
	m.rstore[v] = k
}

// Get value by the k key.
func (m *UniqMap[K, V]) Get(k K) (V, bool) {
	v, ok := m.store[k]
	return v, ok
}

func (m *UniqMap[K, V]) GetByValue(v V) (K, bool) {
	k, ok := m.rstore[v]
	return k, ok
}

