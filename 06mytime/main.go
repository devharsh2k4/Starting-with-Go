package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to study time of Golang!")

	presentTime := time.Now()

	fmt.Println("present time is :", presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 monday"))

	createdDate := time.Date(2023, time.August, 11, 12, 0, 0, 0, time.Local)

	fmt.Println(createdDate)

	fmt.Println(createdDate.Format("01-02-2006 monday"))

}
