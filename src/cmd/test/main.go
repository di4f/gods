package main

import (
	"github.com/surdeus/godat/src/mapx"
	"fmt"
)

func main() {
	m := map[string] string {
		"Key1" : "Value1",
		"Key2" : "Value2",
		"Key3" : "Value3",
	}
	m1 := map[int] string {
		1 : "Val1",
		2 : "Val2",
		7 : "Val7",
	}
	fmt.Println(m)
	fmt.Printf("%q\n", mapx.Keys(m))
	fmt.Printf("%q\n", mapx.Values(m))
	fmt.Printf("%q\n", mapx.Reverse(m))
	fmt.Printf("%v\n", mapx.Reverse(m1))
}
