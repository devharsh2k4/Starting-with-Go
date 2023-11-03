package main

import (
	"fmt"
)

func main() {

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("the address of the ptr is", ptr)
	fmt.Println("The variable whichis stored inside this address", *ptr)

	*ptr = *ptr + 2

	fmt.Println("new value is:", myNumber)
}
