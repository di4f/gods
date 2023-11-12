package stacks

import (
	"github.com/omnipunk/gods"
)

type Stack[V any] interface {
	gods.Container[V]
	Push(V)
	Pop() V
}
