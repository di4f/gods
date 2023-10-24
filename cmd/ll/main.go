package main

import (
	"fmt"
	"github.com/reklesio/gods/lists"
)

func main() {
	list := lists.NewSingly[string]("zero", "one", "two", "three", "four", "five")
	fmt.Println(list)
	list.InsA(0, "after-1", "after-2")
	fmt.Println(list)
	list.InsB(0, "-two", "-one")
	fmt.Println(list)



	//list.Swap(0, 2)
	fmt.Println(list)

	intList := lists.NewSingly[int](100, 5, -1, 1000, 200, 1337)
	fmt.Println(intList)

	//intList.Get(1000)
}
