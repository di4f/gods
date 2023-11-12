package main

import (
	"fmt"
	"github.com/omnipunk/gods/lists"
	"strings"
)

func main() {
	list := lists.NewSingly[string]("zero", "one", "two", "three", "four", "five")
	fmt.Println(list)
	uList := lists.NewSingly(list.Values()...)
	for i, v := range uList.Values() {
		uList.Set(i, strings.ToUpper(v))
	}
	fmt.Println(uList)

	intList := lists.NewSingly[int](100, 5, -1, 1000, 200, 1337)
	fmt.Println(intList)

	//intList.Get(1000)
}
