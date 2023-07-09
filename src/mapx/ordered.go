package mapx

// The type makes the underlying map ordered,
// so every time you pass through all the values
// they will be in the same order.
type OrderedMap[K comparable, V any] struct {
	store map[K] V
	keys []K
}

// Returns the new empty ordered map.
func NewOrdered[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		store: make(map[K] V),
	}
}

// Set or reset the value in the map.
func (m *OrderedMap[K, V]) Set(k K, v V) {
	_, ok := m.store[k]
	if !ok {
		keys = append(keys, k)
	}
	
	m.store[k] = v
}

// Get the value from the map.
func (m *OrderedMap[K, V]) Get(k K) (V, bool) {
	v, ok := m.store[k]
	return v, ok
}

// Return channel of pairs.
func (m *OrderedMap[K, V]) Range() chan Pair

