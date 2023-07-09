package rangex

import (
	"github.com/mojosa-software/godat/src/iterx"
	cnts "golang.org/x/exp/constraints"
)


type Range[V cnts.Ordered] struct {
	start, step, end V
}

// Returns the new Range strucucture with corresponding start, step and end.
// If the values can never reach the end then the function will return nil.
func New[V cnts.Ordered](start, step, end V) *Range[V] {
	return &Range[V]{
		start, step, end,
	}
}

func (r *Range[V]) Chan() iterx.PairChan[int, V] {
	start, step, end := r.start, r.step, r.end
	c := make(iterx.PairChan[int, V])
	go func(){
		var compare func(a, b V) bool
		
		less := func(a, b V) bool {
			return a < b
		}
		
		more := func(a, b V) bool {
			return a > b
		}
		
		if start < end {
			compare = less
		} else {
			compare = more
		}
		
		j := 0
		for i := start ; compare(i, end) ; i += step {
			c <- iterx.Pair[int, V]{
				K:j,
				V:i,
			}
			j++
		}
		
		close(c)
	}()
	
	return c
}

