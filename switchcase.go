//This Go program determines today's date and prints out a specific schedule message based on the day of the month.
//If you run this code today, Monday, 19 October 2026, the program will look at the current day number (19), evaluate it against the schedule rules, and print:
package main

import (
	"fmt"
	"time"
)

func main() {
	timeNow := time.Now()
	var day int = timeNow.Day()
	switch day {
	case 6, 11, 17:
		fmt.Println("Buy Groceries")
	case 5, 16, 27:
		fmt.Println("Play Golf")
	case 3, 19:
		fmt.Println("Go to Church")
	default:
		fmt.Println("Read a magazine")
	}
}
