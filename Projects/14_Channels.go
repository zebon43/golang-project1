package main

import (  
    "fmt"
)

func testchannels(done chan bool) {  
    fmt.Println("Inside Go Routine")
    done <- true //Send data to channel
}
func main() {  
    status := make(chan bool) // Declare a channel variable
    go testchannels(status)  //Invoke the function using go routine
    x:= <-status // Get data from the channel
    fmt.Println("Data from Channel:",x)
	fmt.Println("Main function Ends")
}
