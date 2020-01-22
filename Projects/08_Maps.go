package main

import "fmt"

func main() {
	employeeMap := make(map[int]string) //Short hand Declaration for Map
	employeeMap[1] = "AAAA"
	employeeMap[2] = "BBBB"
	employeeMap[3] = "CCCC"

	//Print values of Map
	fmt.Println("Employee Map Details :", employeeMap)

	employeeMap1 := map[int]string{3: "DDDD", 4: "EEEE", 5: "FFFF"}

	//Print values of Map
	fmt.Println("Employee Map1 Details:", employeeMap1)

	//Access individual element of the Map
	fmt.Println("Employee Details for employee ID 1:", employeeMap[1])

	//Access element of the Map which is not available
	fmt.Println("Employee Details for employee ID 1:", employeeMap[9])

	//Check if the key/value is present or not
	value, ok := employeeMap[2]
	findEmp(value, ok)

	value, ok = employeeMap[9]
	findEmp(value, ok)
	
	//Delete value from map
	delete(employeeMap1,4)
	value, ok = employeeMap1[4]
	findEmp(value, ok)
	fmt.Println("Employee Map1 Details:", employeeMap1)

}

func findEmp(input string, status bool) {
	if status == true {
		fmt.Println("Employee Details for employee ID:", input)
	} else {
		fmt.Println("Employee Details for employee ID not found.")
	}
}
