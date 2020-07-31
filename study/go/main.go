package main

import (
	"./data_structure/other"
	"fmt"
)

func main() {
	fmt.Println("Here is main!")
	//built_in.CommonUsages()
	var stack other.Stack
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Len())
}