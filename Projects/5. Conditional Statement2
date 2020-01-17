package main

import "fmt"

func switchcondition1(in int) {
	switch { //Switch without expression and default case
	case in < 0:
		fmt.Printf("\nThe value %d is Negative.", in)
	case in == 0:
		fmt.Printf("\nThe value %d is Zero.", in)
	case in > 0:
		fmt.Printf("\nThe value %d is Positive.", in)
	}
}

func switchcondition2(in int) {
	switch in % 2 { //Switch with expression and with default case
	case 0:
		fmt.Printf("\nThe value %d is Even.", in)
	case 1:
		fmt.Printf("\nThe value %d is Odd.", in)
	default:
		fmt.Printf("\nThe value %d is not number.", in)
	}
}

func switchcondition3(in int) {
	switch { //Switch with fallthrough
	case in == 0:
		fmt.Printf("\nThe value %d is Zero.", in)
	case in > 0:
		fmt.Printf("\nThe value %d is Positive.", in)
		fallthrough
	case in > 50:
		fmt.Printf("\nThe value %d is between 0 and 50.", in)
	default:
		fmt.Printf("\nThe value %d is Negative number.", in)
	}
}

func main() {

	i1, i2, i3, i4 := -12, 0, 45, 89

	//Check if number if negative or positive
	fmt.Printf("\n------------Neg/Pos------------")
	switchcondition1(i1)
	switchcondition1(i2)
	switchcondition1(i3)
	switchcondition1(i4)

	//Check number is Even or odd
	fmt.Printf("\n\n------------Even/Odd------------")
	switchcondition2(i1)
	switchcondition2(i2)
	switchcondition2(i3)
	switchcondition2(i4)

	//Check if is positive and between 0 and 50 or not
	fmt.Printf("\n\n------------Between 0-50 ------------")
	switchcondition3(i1)
	switchcondition3(i2)
	switchcondition3(i3)
	switchcondition3(i4)

}
