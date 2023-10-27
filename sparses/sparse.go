package sparses
import (
	"github.com/reklesio/gods"
	"cmp"
)

// The general sparse array type.
type Sparse[K cmp.Ordered, V any] interface {
	gods.Container[V]
	// Returns slice of set values.
	Values() []V
	// Returns slice of set keys.
	Keys() []K
	// Set the key to the value.
	Set(K, V)
	// Get the value by the key.
	Get(K) V
	// Returns true if the key is set.
	Has(K) bool
	// Delete the value under the key.
	Del(K)
	// Returns channel of ordered values.
	Chan() chan V
}

