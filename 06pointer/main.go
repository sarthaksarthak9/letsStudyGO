package main

import "fmt"

func main() {
	fmt.Println("welcome lets learn pointers toady")

	// var ptr *int        // generally if we use a var in a funct they create a copy of it insead of modifying actual value
	// fmt.Println("value of pointer is ", ptr)

	myNumber := 23;
	var ptr = &myNumber // whenever & it mean referencing
	fmt.Println("the pointer: ", ptr)  // ptr is the pointer
	fmt.Println("the pointer value: ", *ptr)
	*ptr = *ptr + 2;
	fmt.Println("the new value of myNumber: ", myNumber)

}
