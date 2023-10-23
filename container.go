package gods

// All the containers must implement the interface.
type Container[V any] interface {
	Empty() bool
	Size() int
	Clear()
	Values() []V
	//String() string
}

type Comparator[V any] interface {
	Less(v1, v2 V)
}
