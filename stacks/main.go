package stacks

import (
	"github.com/reklesio/gods"
)

type Stack[V any] interface {
	gods.Container[V]
	Push(V)
	Pop() V
}
