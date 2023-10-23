package main

import (
	"fmt"
	"github.com/reklesio/gods/lists"
)

func main() {
	list := lists.NewSingly[string]("zero", "one", "two", "three", "four", "five")
	fmt.Println(list)
	list.Push("-one", "-two")
	fmt.Println(list)

	list.Swap(0, 2)
	fmt.Println(list)

	intList := lists.NewSingly[int](100, 5, -1, 1000, 200, 1337)
	fmt.Println(intList)

	intList.Sort(func(vi, vj int) bool {
		return vi < vj
	})
	fmt.Println(intList)

	intList.Sort(func(vi, vj int) bool {
		return vj < vi
	})
	fmt.Println(intList)
}
