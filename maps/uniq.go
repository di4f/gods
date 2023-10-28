package maps

// The type describing unique map
// where you can quickly get both
// value by key and vice versa.
type Uniq[K, V comparable] interface {
	Map[K, V]
	GetByValue(V) K
}

type uniqMap[K, V comparable] struct {
	store map[K] V
	rstore map[V] K
}

// The function returns map type where
// you can get, set and delete by value
// since it store everything as ONLY entity
// both for keys and values way.
// Use only when you do not care about the order.
func NewUniq[K, V comparable]() Uniq[K, V] {
	return &uniqMap[K, V]{
		map[K] V{},
		map[V] K{},
	}
}

func (m *uniqMap[K, V]) Empty() bool {
	return len(m.store) == 0
}

func (m *uniqMap[K, V]) Clear() {
	m.store = map[K] V{}
	m.rstore = map[V] K{}
}

func (m *uniqMap[K, V]) Has(k K) bool {
	_, ok := m.store[k]
	return ok
}

func (m *uniqMap[K, V]) Set(k K, v V) {
	m.store[k] = v
	m.rstore[v] = k
}

func (m *uniqMap[K, V]) Del(k K) {
	v := m.store[k]
	delete(m.store, k)
	delete(m.rstore, v)
}

func (m *uniqMap[K, V]) Get(k K) (V) {
	v := m.store[k]
	return v
}

func (m *uniqMap[K, V]) Got(k K) (V, bool) {
	v, ok := m.store[k]
	return v, ok
}

func (m *uniqMap[K, V]) Chan() chan V {
	ret := make(chan V)
	go func() {
		for _, v := range m.store {
			ret <- v
		}
		close(ret)
	}()
	return ret
}

func (m *uniqMap[K, V]) Keys() []K {
	ret := make([]K, len(m.store))
	i := 0
	for k := range m.store {
		ret[i] = k
		i++
	}
	return ret
}

func (m *uniqMap[K, V]) KeyChan() chan K {
	ret := make(chan K)
	go func() {
		for k := range m.store {
			ret <- k
		}
		close(ret)
	}()
	return ret
}

func (m *uniqMap[K, V]) Values() []V {
	ret := make([]V, len(m.store))
	i := 0
	for _, v := range m.store {
		ret[i] = v
		i++
	}
	return ret
}

func (m *uniqMap[K, V]) Size() int {
	return len(m.store)
}

func (m *uniqMap[K, V]) GetByValue(v V) (K) {
	k := m.rstore[v]
	return k
}

