package main

import (
	"fmt"
	"strings"
)

func main() {
	fullName := "james fisk "
	fmt.Println(len(fullName))
	newFullName := strings.TrimRight(fullName, " ")
	fmt.Println(len(newFullName))
	fmt.Println(fullName)
}