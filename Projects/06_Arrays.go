package main

import "fmt"

func main() {
	var arr [3]int
	fmt.Println("Empty Array:", arr)
	fmt.Println("Length of Array arr:", len(arr))

	//Assiging values to Array
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2

	//Printing assigned values
	fmt.Println("\nValues in Integer Array arr:", arr)
	fmt.Println("Length of Array arr:", len(arr))

	arr1 := [4]string{"A", "B", "C", "D"} //Shorthand declaration
	fmt.Println("\nValues in String Array arr1:", arr1)
	fmt.Println("Length of Array arr1:", len(arr1))

	arr2 := [...]float32{0, 1.01, 2.04, 3.06} //Shorthand declaration without length mentioned
	fmt.Println("\nValues in String Array arr2:", arr2)
	fmt.Println("Length of Array arr2:", len(arr2))

	//Looping over array to print values using for
	fmt.Printf("\nPrint values using for loop")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("\nValue at %dth position in array is %d", i, arr[i])
	}
	
	//Looping over array to print values using range (index only)
	fmt.Printf("\n\nPrint values using range (index only)")
	for i := range arr1 {
		fmt.Printf("\nValue at %dth position in array is %s", i, arr1[i])
	}
	
	//Looping over array to print values using range (index & values)
	fmt.Printf("\n\nPrint values using range (index & values)")
	for i, v := range arr2 {
		fmt.Printf("\nValue at %dth position in array is %0.2f", i, v)
	}
}
