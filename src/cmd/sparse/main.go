package main

import (
	"github.com/mojosa-software/godat/src/sparsex"
	"fmt"
)

func main() {
	unord := sparsex.New[int, string](true)
	unord.Set(1, "suck")
	unord.Set(-5, "cock")
	unord.Set(-4, "die")
	unord.Set(-1000, "withme")

	for p := range unord.Chan() {
		fmt.Println(p.K, p.V)
	}

	unord.Sort()
	for p := range unord.Chan() {
		fmt.Println(p.K, p.V)
	}
}
