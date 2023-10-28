package maps

// The type makes the underlying map ordered,
// so every time you pass through all the values
// they will be in the same order.
type orderedMap[K comparable, V any] struct {
	store map[K] V
	keys []K
}

// Returns the new empty ordered map.
func NewOrdered[K comparable, V any]() Map[K, V] {
	return &orderedMap[K, V]{
		store: make(map[K] V),
	}
}

func (m *orderedMap[K, V]) Clear() {
	m.store = map[K]V{}
	m.keys = []K{}
}

func (m *orderedMap[K, V]) Set(k K, v V) {
	_, ok := m.store[k]
	if !ok {
		m.keys = append(m.keys, k)
	}
	
	m.store[k] = v
}

func (m *orderedMap[K, V]) Got(key K) (V, bool) {
	v, ok := m.store[key]
	return v, ok
}

func (m *orderedMap[K, V]) Has(k K) bool {
	_, ok := m.store[k]
	return ok
}

// Get the value from the map.
func (m *orderedMap[K, V]) Get(k K) (V) {
	v := m.store[k]
	return v
}

func (m *orderedMap[K, V]) Del(k K) {
	delete(m.store, k)
	for i, v := range m.keys {
		if v == k {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
		}
	}
}

// Get map keys slice.
func (m *orderedMap[K, V]) Keys() []K {
	return m.keys
}

func (m *orderedMap[K, V]) Values() []V {
	ret := make([]V, len(m.keys))
	i := 0
	for _, k := range m.keys {
		ret[i] = m.store[k]
		i++
	}
	return ret
}

func (m *orderedMap[K, V]) KeyChan() chan K {
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
func (m *orderedMap[K, V]) Chan() chan V {
	chn := make(chan V)
	go func(){
		for _, k := range m.keys {
			chn <- m.Get(k)
		}
		close(chn)
	}()
	
	return chn
}

func (m *orderedMap[K, V]) Empty() bool {
	return len(m.keys) == 0
}

func (m *orderedMap[K, V]) Size() int {
	return len(m.keys)
}

