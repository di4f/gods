package main

import (
	"github.com/mojosa-software/godat/src/poolx"
	
	"fmt"
)

func main() {
	values := []string{
		"zero", "one",
		"should be deleted",
		"two", "three",
	}
	pool := poolx.New[string]()
	for _, v := range values {
		pool.Append(v)
	}
	
	pool.DeleteValue("should be deleted")
	
	for p := range pool.Chan() {
		fmt.Println(p.K, p.V)
	}
}

