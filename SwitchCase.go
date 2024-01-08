package main

import (
	"fmt"
	"time"
)

func TodaysDate() {

	fmt.Printf("Today's ")

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
}
