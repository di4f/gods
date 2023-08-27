package mapx

func Reversed[K, V comparable](m map[K] V) map[V] K {
	r := make(map[V] K)
	for k, v := range m {
		r[v] = k
	}
	return r
}

