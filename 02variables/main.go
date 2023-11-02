package main

import "fmt"

const name = "harsh"

func main() {
	var username string = "Harsh"
	fmt.Println(username)
	fmt.Printf("variable is of type : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type : %T \n", isLoggedIn)

	fmt.Println(name)
}
