package main

import (
	"github.com/mojosa-software/godat/src/llx"
	
	"fmt"
)

func main() {
	ll := llx.New[string]("zero", "one", "two", "three", "four", "five")
	
	ll.Swap(0, 2)
	
	for el := range ll.Chan() {
		fmt.Println(el)
	}
	fmt.Println(ll.Slice())
}

