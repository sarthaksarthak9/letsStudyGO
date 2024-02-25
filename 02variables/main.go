package main

import "fmt"

const logicToken string = "hello world !!" // public

// default value and alias
func main() {
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("variable is of type: %T\n", anothervariable)

	// implicit way
	var siwtch = "hello! I am learning golang"
	fmt.Println(siwtch)

	// no var style

	noOfusers := 3000
	fmt.Println(noOfusers)

	fmt.Println(logicToken)
	fmt.Printf("variable is of type: %T\n", logicToken)

}
