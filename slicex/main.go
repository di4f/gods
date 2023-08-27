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

func ConvStr[V1, V2 ~string](v1 []V1) []V2 {
	ret := []V2{}
	for i := range v1 {
		ret = append(ret, V2(v1[i]))
	}

	return ret
}

func ConvInt[V1, V2 ~int](v1 []V1) []V2 {
	ret := []V2{}
	for i := range v1 {
		ret = append(ret, V2(v1[i]))
	}

	return ret
}

