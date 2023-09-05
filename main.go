package main

import (
	list "dataStructures/list"
	"fmt"
)

func main() {
	fmt.Println("hi")
	hi := list.NewList[int]()
	hi.Add(1)
	hi.Add(2)

}
