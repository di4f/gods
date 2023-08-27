package iterx

func ChanToSlice[V any](c chan V) []V {
	ret := []V{}
	for v := range c {
		ret = append(ret, v)
	}
	
	return ret
}

