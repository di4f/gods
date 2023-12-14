package stacks

import (
	"github.com/di4f/gods"
)

type Stack[V any] interface {
	gods.Container[V]
	Push(V)
	Pop() V
}
