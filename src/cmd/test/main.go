package main

import (
	"github.com/surdeus/godat/src/mapx"
	"github.com/surdeus/godat/src/slicex"
	"github.com/surdeus/godat/src/llx"
	"fmt"
)

type Struct struct {
	Name string
	Value int
}

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
	s := []Struct {
		{"Name1", 1},
		{"Name2", 2},
	}

	fmt.Println(m)
	fmt.Println(slicex.MakeMap(
		s,
		func(s []Struct, i int) string {
			return s[i].Name
		},
	))

	fmt.Printf("%q\n", mapx.Keys(m))
	fmt.Printf("%q\n", mapx.Values(m))
	fmt.Printf("%q\n", mapx.Reverse(m))
	fmt.Printf("%v\n", mapx.Reverse(m1))
	
	ll := llx.New[int]()
	ll.Append(0)
	ll.Append(1)
	ll.Append(2)
	ll.Set(1, 256)
	for p := range ll.Range() {
		fmt.Println(p)
	}
}
