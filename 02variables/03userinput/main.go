package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hey")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our app")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
}
