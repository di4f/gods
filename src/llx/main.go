package llx

// Linked list X .
// The package implements better variation of
// linked list than in standard library since it uses
// the new conception of generics.

type LinkedList[V any] struct {
	// First empty element (not used to store values).
	// For fast pushing.
	before *Element[V]
	// Points to the last for fast appending.
	last *Element[V]
	// Length.
	ln int
}

type Element[V any] struct {
	next *Element[V]
	value V
}

type Pair[V any] struct {
	I int
	V V
}


// Returns new empty linked list storing the V type.
func New[V any]() *LinkedList[V] {
	return &LinkedList[V]{
		&Element[V]{},
		nil,
		0,
	}
}

// Get length of the linked list.
func (ll *LinkedList[V]) Len() int {
	return ll.ln
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

func (ll *LinkedList[V]) Del(i int) (bool) {
	if ll.ln <= i {
		return false
	}
	
	if i == 0 {
		ll.before.next =
			ll.before.next.next
		ll.ln--
		return true
	}
	
	el1, _ := ll.GetEl(i-1)
	if i == ll.ln - 1 {
		el1.next = nil
	} else {
		el2, _ := ll.GetEl(i+1)
		el1.next = el2
	}
	
	ll.ln--
	return true
}

// Push in the beginning of the list.
func (ll *LinkedList[V]) Push(v V) {
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
func (ll *LinkedList[V]) Append(v V) {
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

func (ll *LinkedList[V]) First() *Element[V] {
	return ll.before.next
}

func (ll *Element[V]) Value() V {
	return ll.value
}

func (ll *Element[V]) Next() *Element[V] {
	return ll.next
}

func (ll *LinkedList[V]) Last() *Element[V] {
	return ll.last
}

// Returns a channel of Pair that contains index and the value.
func (ll *LinkedList[V]) Range() chan Pair[V] {
	chn := make(chan Pair[V])
	go func(){
		i := -1
		el := ll.before
		for el.next != nil {
			i++
			el = el.next
			chn <- Pair[V]{i, el.value}
		}
		close(chn)
	}()
	return chn
}

