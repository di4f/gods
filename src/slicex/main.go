package slicex

func MakeMap[K comparable, V any](
	values []V,
	fn func([]V, int) (K),
) map[K] V {
	var k K

	r := make(map[K] V)
	for i, _ := range values {
		k = fn(values, i)
		r[k] = values[i]
	}

	return r
}
