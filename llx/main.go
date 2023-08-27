package llx

// Linked list X .
// The package implements better variation of
// linked list than in standard library since it uses
// the new conception of generics.

// The type represents linked list data structure.
type LinkedList[V any] struct {
	// First empty element (not used to store values).
	// For fast pushing.
	before *Element[V]
	// Points to the last for fast appending.
	last *Element[V]
	// Length.
	ln int
}

// The type represents element of the linked list.
type Element[V any] struct {
	next *Element[V]
	value V
}

// Returns new empty linked list storing the V type.
func New[V any](values ...V) *LinkedList[V] {
	ret := &LinkedList[V]{
		&Element[V]{},
		nil,
		0,
	}

	ret.Append(values...)
	return ret
}

// Get length of the linked list.
func (ll *LinkedList[V]) Length() int {
	return ll.ln
}

func (ll *LinkedList[V]) Len() int {
		return ll.Length()
}

// Get the index-indexed element itself.
func (ll *LinkedList[V]) GetEl(index int) (*Element[V], bool) {
	if ll.ln <= index {
		return nil, false
	}
	p := ll.before
	for i := 0 ; i <= index ; i++ {
		p = p.next
	}
	
	return p, true
}

// Get the value of index-indexed element.
func (ll *LinkedList[V]) Get(index int) (V, bool) {
	el, ok := ll.GetEl(index)
	var v V
	if ok {
		v = el.value
	}
	
	return v, ok
}

// Set the new value in i-indexed element.
func (ll *LinkedList[V]) Set(i int, v V) (bool) {
	el, ok := ll.GetEl(i)
	if !ok {
		return false
	}
	
	el.value = v
	return true
}

// Insert the V value before the i-th element.
func (ll *LinkedList[V]) InsertBefore(v V, i int) {
		if i == 0 {
			ll.before = &Element[V]{
				value: v,
				next: ll.before.next,
			}
			return
		}
		el, ok := ll.GetEl(i-1)
		if !ok {
			panic("index out of range")
		}

		el.next = &Element[V]{
			value: v,
			next: el.next,
		}
}

// Insert the V value after the i-th element.
func (ll *LinkedList[V]) InsertAfter(i int, v V) {
		el, ok := ll.GetEl(i)
		if !ok {
			panic("index out of range")
		}

		el.next = &Element[V]{
			value: v,
			next: el.next,
		}
}

// Swap element values indexed by i1 and i2.
// Panic on "index out of range".
func (ll *LinkedList[V]) Swap(i1, i2 int) {
	if i1 == i2 {
		return
	}
	
	max := ll.ln - 1
	if i1 < 0 || i2 < 0 || i1 > max || i2 > max {
		panic("index out of range")
	}
	
	el1, _ := ll.GetEl(i1)
	el2, _ := ll.GetEl(i2)
	
	el1.value, el2.value =
		el2.value, el1.value
}

// Deletes the element by its index.
func (ll *LinkedList[V]) Delete(i int) {
	if ll.ln <= i && i < 0 {
			panic("index out of range")
	}
	
	if i == 0 {
		ll.before.next =
			ll.before.next.next
		ll.ln--
		return
	}
	
	el1, _ := ll.GetEl(i-1)
	if i == ll.ln - 1 {
		el1.next = nil
	} else {
		el2, _ := ll.GetEl(i+1)
		el1.next = el2
	}
	
	ll.ln--
}

// Alias to Delete method.
func (ll *LinkedList[V]) Del(i int) {
	ll.Delete(i)
}

// Push in the beginning of the list.
func (ll *LinkedList[V]) Push(values ...V) {
	for _, value := range values {
		ll.push(value)
	}
}

// Push in the beginning of the list.
func (ll *LinkedList[V]) push(v V) {
	prevNext := ll.before.next
	nextNext := &Element[V]{
		next: prevNext,
		value: v,
	}
	ll.before.next = nextNext
	
	ll.ln++
	if ll.ln == 1 {
		ll.last = ll.before.next
	}
}

// Append to the end of the list.
func (ll *LinkedList[V]) Append(values ...V) {
	for _, value := range values {
		ll.gappend(value)
	}
}

func (ll *LinkedList[V]) gappend(v V) {
	if ll.ln == 0 {
		ll.Push(v)
		return
	}
	
	last := &Element[V]{
		next: nil,
		value: v,
	}
	
	lastBuf := ll.last
	lastBuf.next = last
	ll.last = last
	
	ll.ln++
}

// Returns the first element of the linked list.
func (ll *LinkedList[V]) First() *Element[V] {
	return ll.before.next
}

// Get elements value.
func (ll *Element[V]) Value() V {
	return ll.value
}

// Returns the next element. If the returned value == nil,
// then it is the last element.
func (ll *Element[V]) Next() *Element[V] {
	return ll.next
}

// Returns the last element.
func (ll *LinkedList[V]) Last() *Element[V] {
	return ll.last
}

// Returns a channel with values ordered as in list.
func (ll *LinkedList[V]) Chan() chan V {
	chn := make(chan V)
	go func(){
		el := ll.before
		for el.next != nil {
			el = el.next
			chn <- el.value
		}
		close(chn)
	}()
	return chn
}

// Returns slice of values in the list ordered as in the list.
func (ll *LinkedList[V]) Slice() []V {
	buf := make([]V, ll.Length())
	chn := ll.Chan()

	i := 0
	for v := range chn {
		buf[i] = v
		i++
	}

	return buf
}

func (ll *LinkedList[V]) Sort() {
}

