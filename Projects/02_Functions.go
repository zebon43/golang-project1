package main

import (
	"fmt"
)

func add(x, y int) {
	add_Result := x + y
	fmt.Println("Addition is:", add_Result)
}

func sub(x, y int) int {
	sub_Result := x - y
	return sub_Result
}

func calculate(x, y int) (int, float32) {
	div := float32(x) / float32(y)
	mul := x * y
	return mul, div
}

func mod(x, y int) (result int) {
	result = x % y
	return
}

func main() {

	//Decalre 2 intergers
	var i1, i2 = 20, 12

	// Function Calls
	add(i1, i2)               // Call the function only
	
	sub_Result := sub(i1, i2) // Call the fucntion and store the values
	fmt.Println("Subtraction is:", sub_Result)
	
	mul_Result, div_Result := calculate(i1, i2) // Call function with multiple returns
	fmt.Printf("\nMultiplication is: %d \nDivision is: %f", mul_Result, div_Result)
	
	fmt.Println("\nAddition is:", mod(i1, i2)) //Directly call function and print
}
