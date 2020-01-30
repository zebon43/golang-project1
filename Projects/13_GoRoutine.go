package main

import (
	"fmt"
	"time"
)

//Define Function to print Seven numbers with a delay of 50 Milliseconds
func goroutine1() {
	for i := 1; i <= 7; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Println("I am Go Routine 1 - ", i)
	}
}

//Define Function to print Four numbers witha  delay of 90 Milliseconds
func goroutine2() {
	for i := 1; i <= 4; i++ {
		time.Sleep(90 * time.Millisecond)
		fmt.Println("I am Go Routine 2 - ", i)
	}
}

func main() {
	fmt.Println("Main Program Starts.")
	go goroutine1() //Call Function 1 with go call
	go goroutine2() //Call Function 2 with go call
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("Main Program Ends.")
}
