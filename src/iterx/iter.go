package iterx

// Implementing the interface lets us iterate through the
// the data by lightweight channels.
type Channeler[K any, V any] interface {
	Chan() PairChan[K, V]
}

// Implementing the interface provides the way to
// convert the type to slice.
type Slicer[V any] interface {
	Slice() []V
}

// Implementing the interface provides us the way to 
// convert the type to map.
type Mapper[K comparable, V any] interface {
	Map() map[K] V
}

