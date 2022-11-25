package slicex

func MakeMap[K comparable, V any](
	values []V,
	fn func(V) (K),
) map[K] V {
	var k K

	r := make(map[K] V)
	for _, v := range values {
		k = fn(v)
		r[k] = v
	}

	return r
}
