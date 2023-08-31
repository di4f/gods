package main

import (
	"github.com/mojosa-software/godat/mapx"
	"fmt"
)

func main() {
	ord := mapx.NewOrdered[int, string]()
	ord.Set(1, "one")
	ord.Set(5, "five")
	ord.Set(2, "two")
	ord.Del(5)
	for v := range ord.Chan() {
		fmt.Println(v)
	}
}
