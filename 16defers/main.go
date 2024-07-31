// defer - when a funct execute it execute line by line. But when we use defer keywords the next line going to
// execute at the end of the function. If we use multiple defer keywords the order of execution is LIFO

package main

import "fmt"


func main(){
	defer fmt.Println("World")
	defer fmt.Println("three")
	defer fmt.Println("two")
	defer fmt.Println("one")
	fmt.Println("Hello")
	myDefer()
}

func myDefer(){
	for i := 0; i < 5; i++{
		defer fmt.Println(i)
	} 
}
