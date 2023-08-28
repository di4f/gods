package main

import (
	"github.com/mojosa-software/godat/sparsex"
	"fmt"
)

func main() {
	arr := sparsex.New[float64, string](map[float64] string {
			5: "something at 5",
			12: "new shit",
			50: "die",
		}).
		Default("<NIL>")

	arr.Del(12)
	for v := range arr.Chan() {
		fmt.Println(v)
	}
}
