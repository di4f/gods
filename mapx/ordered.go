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
		m.keys = append(m.keys, k)
	}
	
	m.store[k] = v
}

func (m *OrderedMap[K, V]) Has(k K) bool {
	_, ok := m.store[k]
	return ok
}

// Get the value from the map.
func (m *OrderedMap[K, V]) Get(k K) (V) {
	v := m.store[k]
	return v
}

func (m *OrderedMap[K, V]) Del(k K) {
	delete(m.store, k)
	for i, v := range m.keys {
		if v == k {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
		}
	}
}

// Get map keys slice.
func (m *OrderedMap[K, V]) Keys() []K {
	return m.keys
}

func (m *OrderedMap[K, V]) KeyChan() chan K {
	chn := make(chan K)
	go func() {
		for _, v := range m.keys {
			chn <- v
		}
		close(chn)
	}()
	return chn
}

// Return channel of ordered values.
func (m *OrderedMap[K, V]) Chan() chan V {
	chn := make(chan V)
	go func(){
		for _, k := range m.keys {
			chn <- m.Get(k)
		}
		close(chn)
	}()
	
	return chn
}

