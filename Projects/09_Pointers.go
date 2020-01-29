package main

import "fmt"

func main() {
	var x int = 1502

	//Decalre a pointer
	var ptr_x *int

	//Print default value of pointer
	fmt.Println("Default value of pointer is:", ptr_x)

	//Assign Addres of x to pointer
	ptr_x = &x

	//Printe the pointer type and its pointer
	fmt.Printf("Type of pointer variable ptr_x is %T\n", ptr_x)
	fmt.Println("Address of integer variable x is ", ptr_x)
	fmt.Printf("Value of pointer variable ptr_x after assignment is %v", *ptr_x)

	//Decalre a pointer using new
	ptr_y := new(string)
	fmt.Printf("\n\nType of pointer variable ptr_y is %T", ptr_y)
	fmt.Printf("\nValue of pointer variable ptr_y is %v", *ptr_y)
	fmt.Printf("\nAddress of pointer variable ptr_y is %v", ptr_y)

	*ptr_y = "New Value"
	fmt.Printf("\nValue of pointer variable ptr_y after assignement is: %v", *ptr_y)

	//Function and pointers - Pass by reference
	changeValue(ptr_x)
	//Print the value of variable to check if the value has been changed or not
	fmt.Printf("\n\nValue of x after function call is %v", x)

}

func changeValue(input *int) {
	*input++
}
