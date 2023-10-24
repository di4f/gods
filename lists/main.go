package lists

import (
	"github.com/reklesio/gods"
	"github.com/reklesio/gods/stacks"
)

// The interface all the lists must implement.
type List[V any] interface {
	gods.Container[V]
	stacks.Stack[V]

	// Like push in stacks.
	Put(...V)
	// Get length of the list.
	Len() int
	// Get the value by index.
	Get(int) V
	// Delete the value by index.
	Del(int)
	// Change already existing value.
	Set(int, V)
	// Add the values to the end of the list.
	Add(...V)

	// Insert the value before the specified index.
	InsB(int, ...V)
	// Ansert values after the specified index.
	InsA(int, ...V)

	// Swap elements by indexes specified in arguments.
	Swap(i, j int)

	// Return channel with all the values.
	Chan() chan V

	// The sort function that gets the Less function as argument
	// and sorts the list corresponding to it.
	Sort(gods.LessFunc[V])
}

