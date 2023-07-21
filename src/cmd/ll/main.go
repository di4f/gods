package main

import (
	"github.com/mojosa-software/godat/src/llx"
	
	"fmt"
)

func main() {
	ll := llx.New[string]()
	ll.Append("zero")
	ll.Append("one")
	ll.Append("two")
	ll.Append("three")
	ll.Append("four")
	ll.Push("minus one")
	
	ll.Swap(1, 3)
	
	for p := range ll.Chan() {
		fmt.Println(p.K, p.V)
	}
}

