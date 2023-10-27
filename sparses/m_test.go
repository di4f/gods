package sparses

import (
	"testing"
)


func TestNew(t *testing.T) {
	s := New[int, string](0, nil)
	if aVal := s.Empty() ; if aVal != true {
		t.Errorf("Got %v expected %v", aVal, )
	}

	s = New[int, string]
}
