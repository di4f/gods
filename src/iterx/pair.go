package iterx

// The type describes pair of key and value.
type Pair[K any, V any] struct {
	V V
	K K
}

// The type describes channel of pairs.
type PairChan[K any, V any] chan Pair[K, V]

// Slice of pairs.
type Pairs[K any, V any] []Pair[K, V]

// ForEach for channels, like in JS.
// Be careful since the function does not close the
// channel so if fn breaks the loop then there can
// be values left.
func (pc PairChan[K, V]) ForEach(fn func(k K, v V) bool) {
	for p := range pc {
		if !fn(p.K, p.V) {
			break
		}
	}
}


