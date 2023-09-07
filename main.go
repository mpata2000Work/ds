package main

import (
	list "dataStructures/list"
	"fmt"
)

func comp(v1 int, v2 int) int {
	return v1 - v2
}

func main() {
	fmt.Println("hi")
	hi := list.NewLinkedList[int](comp)
	hi.AddLast(0)
	hi.AddLast(1)
	hi.AddLast(2)
	hi.PrettyPrint()
	hi.AddLast(3)
	hi.AddLast(4)
	hi.PrettyPrint()

	fmt.Println("Size: ", hi.Size())
	r, _ := hi.Get(1)
	fmt.Println("At index 1: ", r)
	r, _ = hi.Get(4)
	fmt.Println("At index 4: ", r)
	t, err := hi.Get(5)
	fmt.Println("At index 5: ", t)
	fmt.Println("Err: ", err)

}
