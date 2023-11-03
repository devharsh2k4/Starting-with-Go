package main

import "fmt"

func main() {
	fmt.Println("Arrays")

	var fruitList [4]string

	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[3] = "guava"

	fmt.Println(fruitList)
	fmt.Println(len(fruitList))

	var veggies = [3]string{"potato,cabbage,cauliflower"}
	fmt.Println(veggies)
}
