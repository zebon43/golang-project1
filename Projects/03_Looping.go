package main

import "fmt"

func fibo(input int) {
	n1, n2, temp := 0, 1, 0
	fmt.Printf("Fiboacci : %d %d ", n1, n2)
	for i := 0; i < input; i++ {
		temp = n1 + n2
		fmt.Printf("%d ", temp)
		n1 = n2
		n2 = temp
	}
}

func fact(input int) (output int) {
	for i := 0; i < input; i++ {
		if input == 1 {
			return 1
		}
		output = input * fact(input-1)
	}
	return
}

func breakloop(input int) {
	for i := 0; i < input; i++ {

		if i == 5 {
			break //Exit the loop if i is 5
		}
		fmt.Printf("%d ", i)
	}
}

func continueloop(input int) {
	for i := 0; i < input; i++ {

		if i == 5 {
			continue //Exit the loop if i is 5
		}
		fmt.Printf("%d ", i)
	}
}

func infinteloop() {
	for {
		//Infinte loop as the Init, condition and Post are optional
	}
}

func main() {
	input := 10

	fibo(input) //print the fibo series

	fmt.Println("\nFactorial:", fact(input)) // print factorial

	fmt.Println("\nBreak Loop:")
	breakloop(input) //print 1 to input with break at 5

	fmt.Println("\nContinue loop:")
	continueloop(input) //print 1 to input with skipping 5
}
