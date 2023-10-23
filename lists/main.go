package lists

import (
	"github.com/reklesio/gods"
)

// The interface all the lists must implement.
type List[V any] interface {
	gods.Container[V]
	Push(...V)
	// Get length of the list
	Len() int
	// Get the value by index.
	Get(int) V
	// Delete the value by index.
	Del(int)
	// Change already existing value.
	Set(int, V)
	// Add the values
	Add(...V)

	// Insert the value before the specifed index.
	InsB(V, int)
	// Insert the value after the specified index.
	InsA(int, V)

	// Swap elements by indexes specified in arguments.
	Swap(i, j int)

	// The sort function that gets the Less function as argument
	// and sorts the list corresponding to it.
	Sort(gods.LessFunc[V])
}

