package main

import (
	"fmt"
)

func main() {

	var x int = -10

	//if only
	if x < 0 {
		fmt.Printf("\nThe number %d is negative.", x)
	}

	x = 10
	//if and else
	if x < 0 {
		fmt.Printf("\nThe number %d is negative.", x)
	} else {
		fmt.Printf("\nThe number %d is positive.", x)
	}

	//if and else with initialized statement
	if x := 5; x < 0 {
		fmt.Printf("\nThe number %d is negative.", x)
	} else {
		fmt.Printf("\nThe number %d is positive.", x)
	}

	//Nested if and else
	if x := 0; x < 0 {
		fmt.Printf("\nThe number %d is negative.", x)
	} else {
		if x == 0 {
			fmt.Printf("\nThe number %d is neither positive nor negative.", x)
		} else {
			fmt.Printf("\nThe number %d is positive.", x)
		}
	}
}
