package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	phone     int
	address   string
}

func printPerson(temp Person) {
	//Print the details of person
	fmt.Println("Person Details:", temp)
}

func printPersonDetails(temp Person) {
	//Print the details of person
	fmt.Println("First Name:", temp.firstName)
	fmt.Println("First Name:", temp.lastName)
	fmt.Println("First Name:", temp.age)
	fmt.Println("First Name:", temp.phone)
	fmt.Println("First Name:", temp.address)
}

type Address struct {
	building, street, city, state, country string
	pincode                                int
}

type PersonNew struct {
	firstName string
	lastName  string
	age       int
	phone     int
	address   Address
}

func main() {

	//Create a variable of type Person
	var p1 Person

	//Print default values which will be 0 or Nil
	fmt.Println("------ Person Details Default Values ------")
	printPerson(p1)
	printPersonDetails(p1)

	//Assign values
	p1 = Person{
		firstName: "AAAA",
		lastName:  "BBBB",
		age:       12,
		phone:     1234567890,
		address:   "Building Number, Street, City, Country, Pin Code",
	}

	fmt.Println("\n------ Person Details assigned Values ------")
	printPerson(p1)
	printPersonDetails(p1)

	//Anonymous Struct - Direct variable creation with strcuture
	h1 := struct{ firstname, lastname string }{firstname: "FFFF", lastname: "LLLL"}

	//Print Values
	fmt.Println("\n\n------ Human Details Default Values ------")
	fmt.Println("Human Details:", h1)

	//Creating variable for Nested Struct
	p2 := PersonNew{
		firstName: "AAAA",
		lastName:  "BBBB",
		age:       12,
		phone:     1234567890,
		address: Address{
		"Building number 1",
		"Street 10",
		"City 100",
		"Country 1000",
		"State 10000",
		1000000,},
	}
	
	fmt.Println("\n\n------ Nested Struct Values ------")
	fmt.Println("Person Details:", p2)

}
