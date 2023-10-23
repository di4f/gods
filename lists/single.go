package lists

import (
	"github.com/reklesio/gods"
	"sort"
	"fmt"
)

// Linked list X .
// The package implements better variation of
// linked list than in standard library since it uses
// the new conception of generics.

// The type represents singly linked list data structure.
type sLinkedList[V any] struct {
	// First empty element (not used to store values).
	// For fast pushing.
	before *sElement[V]
	// Points to the last for fast appending.
	last *sElement[V]
	// Length.
	ln int
}

// The type represents element of the linked list.
type sElement[V any] struct {
	next *sElement[V]
	value V
}

func newSingly[V any](values ...V) *sLinkedList[V] {
	ret := &sLinkedList[V]{
		before: &sElement[V]{},
		last: nil,
		ln: 0,
	}

	ret.Add(values...)
	return ret
}

// Returns new empty linked list storing the V type.
func NewSingly[V any](values ...V) List[V] {
	return newSingly[V](values...)
}

func (ll *sLinkedList[V]) Empty() bool {
	return ll == nil
}

func (ll *sLinkedList[V]) Size() int {
	return ll.ln
}

func (ll *sLinkedList[V]) Clear() {
	buf := newSingly[V]()
	*ll = *buf
}

func (ll *sLinkedList[V]) Len() int {
		return ll.ln
}


// Get the index-indexed element itself.
func (ll *sLinkedList[V]) getEl(index int) *sElement[V] {
	if ll.ln <= index || index < 0 {
		panic(gods.IndexRangeErr)
	}
	p := ll.before
	for i := 0 ; i <= index ; i++ {
		p = p.next
	}
	
	return p
}

// Get the value of index-indexed element.
func (ll *sLinkedList[V]) Get(index int) V {
	return ll.getEl(index).value
}

// Set the new value in i-indexed element.
func (ll *sLinkedList[V]) Set(i int, v V) {
	el := ll.getEl(i)
	el.value = v
}

// Insert the V value before the i-th element.
func (ll *sLinkedList[V]) InsB(v V, i int) {
		if i == 0 {
			ll.before = &sElement[V]{
				value: v,
				next: ll.before.next,
			}
			return
		}
		el := ll.getEl(i-1)
		el.next = &sElement[V]{
			value: v,
			next: el.next,
		}
}

// Insert the V value after the i-th element.
func (ll *sLinkedList[V]) InsA(i int, v V) {
		el := ll.getEl(i)
		el.next = &sElement[V]{
			value: v,
			next: el.next,
		}
}

// Swap element values indexed by i1 and i2.
// Panic on "index out of range".
func (ll *sLinkedList[V]) Swap(i1, i2 int) {
	if i1 == i2 {
		return
	}
	
	el1 := ll.getEl(i1)
	el2 := ll.getEl(i2)
	
	el1.value, el2.value =
		el2.value, el1.value
}

// Deletes the element by its index.
func (ll *sLinkedList[V]) Del(i int) {
	if i == 0 {
		ll.before.next =
			ll.before.next.next
		ll.ln--
		return
	}
	
	el1 := ll.getEl(i-1)
	if i == ll.ln - 1 {
		el1.next = nil
	} else {
		el2 := ll.getEl(i+1)
		el1.next = el2
	}
	
	ll.ln--
}

// Push in the beginning of the list.
func (ll *sLinkedList[V]) Push(values ...V) {
	for _, value := range values {
		ll.push(value)
	}
}

// Push in the beginning of the list.
func (ll *sLinkedList[V]) push(v V) {
	prevNext := ll.before.next
	nextNext := &sElement[V]{
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
func (ll *sLinkedList[V]) Add(values ...V) {
	for _, value := range values {
		ll.gappend(value)
	}
}

func (ll *sLinkedList[V]) gappend(v V) {
	if ll.ln == 0 {
		ll.Push(v)
		return
	}
	
	last := &sElement[V]{
		next: nil,
		value: v,
	}
	
	lastBuf := ll.last
	lastBuf.next = last
	ll.last = last
	
	ll.ln++
}

// Returns the first element of the linked list.
func (ll *sLinkedList[V]) First() *sElement[V] {
	return ll.before.next
}

// Get elements value.
func (ll *sElement[V]) Value() V {
	return ll.value
}

// Returns the next element. If the returned value == nil,
// then it is the last element.
func (ll *sElement[V]) Next() *sElement[V] {
	return ll.next
}

// Returns the last element.
func (ll *sLinkedList[V]) Last() *sElement[V] {
	return ll.last
}

// Returns a channel with values ordered as in list.
func (ll *sLinkedList[V]) Chan() chan V {
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
func (ll *sLinkedList[V]) Values() []V {
	buf := make([]V, ll.Len())
	i := 0
	el := ll.before
	for el.next != nil {
		el = el.next
		buf[i] = el.value
		i++
	}

	return buf
}

func (ll *sLinkedList[V]) String() string {
	return fmt.Sprintf("%v", ll.Values())
}

func (ll *sLinkedList[V]) Sort(fn gods.LessFunc[V]) {
	sort.Sort(gods.CustomSort[V]{
		CustomSorter: ll,
		LessFunc: fn,
	})
}

