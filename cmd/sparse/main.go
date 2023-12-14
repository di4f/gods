package main

import (
	"github.com/di4f/gods/maps"
	"fmt"
)

func main() {
	arr := maps.NewSparse[float32, string]("default", map[float32]string{
		5:  "something at 5",
		12: "new shit 12",
		50: "die 50",
	})
	for i := 0; i <= 50; i++ {
		fmt.Println(arr.Get(float32(i)))
	}
	fmt.Println(arr.Size())

	arr.Del(5)
	arr.Del(12)
	arr.Del(50)
	for i := 0; i <= 50; i++ {
		fmt.Println(arr.Get(float32(i)))
	}
	fmt.Println(arr.Size())

	fmt.Printf("%v\n", arr)
	arr.Del(12)
	arr.Set(10, "at 10")
	arr.Set(100, "at 100")
	arr.Set(49, "at 100")
	arr.Set(48.5, "at 100")
	fmt.Printf("%v\n", arr)
}
