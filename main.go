package main

import (
	list "dataStructures/list"
	"fmt"
)

func main() {
	fmt.Println("hi")
	hi := list.NewList[int]()
	hi.AddLast(0)
	hi.AddLast(1)
	hi.AddLast(2)
	hi.PreatyPrint()
	hi.AddLast(3)
	hi.AddLast(4)
	hi.PreatyPrint()

	fmt.Println("Size: ", hi.Size())
	fmt.Println("Head ", hi.Peek())
	fmt.Println("Tail ", hi.Top())
	r, _ := hi.GeatAt(1)
	fmt.Println("At index 1: ", r)
	r, _ = hi.GeatAt(4)
	fmt.Println("At index 4: ", r)
	t, err := hi.GeatAt(5)
	fmt.Println("At index 5: ", t)
	fmt.Println("Err: ", err)

}
