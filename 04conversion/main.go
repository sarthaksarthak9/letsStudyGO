package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the pizza app")
	fmt.Println("Please rate the pizza from 1 to 5")

	reader := bufio.NewReader(os.Stdin) // it always give a string

	input, _ := reader.ReadString('\n') // till upto new lines hit
	fmt.Println("Thank for the rating", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// strconv return two values that's why we have use numRating and err
	// strings.TrimSpace() to trimming up the space that is coming up by hitting enter
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("we have added:", numRating+1)
	}
}
