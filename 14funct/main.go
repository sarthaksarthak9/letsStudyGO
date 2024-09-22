package main

import "fmt"

func main() {
	greeter()
	fmt.Println("welcome to functions in go")

	result := adder(1, 2
	fmt.Println("The result is ", result)

	proResult, myMess := proAdder(2,3,4,5,6) // or we can proResult, _
	fmt.Println("The proRes is", proResult)
	fmt.Println("The message", myMess)
}
 
func proAdder(values ...int) (int, string) {  // ... is a varietic func it can accept any type value
	total := 0;

	for _, val := range values{
		total += val
	}
	return total, "This is a pro result"
}

func greeter(){
	fmt.Println("namaste!")
}

func adder(a int, b int) int {
	return a + b
}

// anonymous function do exit 

// func (){
// 	fmt.Println("I am an anonymous function")
// }()



