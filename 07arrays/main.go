package main

import "fmt"

// In golang we don't use array so freq compared to other language
// it don;t having sort and other stuff

func main() {
	fmt.Println("Welcome lets learn array in golang")

	var fruits [4]string

	fruits[0] = "apple"
	fruits[1] = "mango"
	fruits[3] = "chicku"

	fmt.Println(fruits)  // u can observer the uneven space in the output
	fmt.Println(len(fruits))

	var vegList = [3] string{"potato", "beans", "mushroom"}
}
