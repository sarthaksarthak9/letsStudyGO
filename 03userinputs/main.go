package main

import (
	"fmt"
)

func main() {
	welcome := "welcome to the user input"
	fmt.Println(welcome)

	// techWorldwithNana.

	// var username string
	// fmt.Println("Type ur name :) ")
	// fmt.Scan(&username)

	// fmt.Printf("your name is %v", username)
	// fmt.Println("")

	reader := bufio.NewReader(os.Stdin) // This line creates a new bufio.Reader that reads from the standard input (os.Stdin).
	fmt.Println("enter the rating of our pizza: ")

	// comma ok || comma error
	// { in go we don't have error catch, the paradigm of go expect to treat error like a variable, like true false }

	input, _ := reader.ReadString('\n') // all the input goes into the input variable, we here don't care much of the error so _

	fmt.Println("Thanks for rating", input) // when we care of error then _, err , the err goes to err variable
	fmt.Printf("Type of this rating is %T", input)
	fmt.Println("")

}
