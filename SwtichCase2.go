package main

import (
	"fmt"
	"time"
)

func DaystoToday() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Printf("This is the day")
	case today + 1:
		fmt.Printf("One day to go")
	case today + 2:
		fmt.Printf("You must wait two days yet")
	default:
		fmt.Printf("Too far away")
	}
}
