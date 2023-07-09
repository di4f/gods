package mapx

// The package implements some more specific
// map structures for special uses.

func Keys[K comparable, V any](m map[K] V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}

	return r
}

func Values[K comparable, V any](m map[K] V) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}

	return r
}

func Reverse[K, V comparable](m map[K] V) map[V] K {
	r := make(map[V] K)
	for k, v := range m {
		r[v] = k
	}

	return r
}

