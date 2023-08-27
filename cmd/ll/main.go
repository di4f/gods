package main

import (
	"fmt"
	"github.com/mojosa-software/godat/llx"
)

func main() {
	ll := llx.New[string]("zero", "one", "two", "three", "four", "five")
	ll.Push("-one", "-two")

	ll.Swap(0, 2)
	fmt.Println(ll.Slice())
}
