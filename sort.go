package gods

type Getter[V any] interface {
	Get(int) V
}

type Swapper interface {
	Swap(i, j int)
	Len() int
}

type CustomSorter[V any] interface {
	Getter[V]
	Swapper
}

type LessFunc[V any] func(v1, v2 V) bool
// The type implements way to sort
// swappers via custom function.
type CustomSort[V any] struct {
	CustomSorter[V]
	LessFunc LessFunc[V]
}

func (cs CustomSort[V]) Less(i, j int) bool {
	vi, vj := cs.Get(i), cs.Get(j)
	return cs.LessFunc(vi, vj)
}

