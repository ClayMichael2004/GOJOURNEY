package main

import "fmt"

// JVAL is global coz its declared outside the main function
const JVAL = 5.87

func main() {
	//PRO and RIGHT are local hence can only be used or accessed inside this function
	const PRO = "Bike has wheels"
	fmt.Println("Cycle", PRO)
	fmt.Println("JValue is", JVAL)
	const RIGHT = true
	fmt.Println("Checking Boolean", RIGHT)
}
