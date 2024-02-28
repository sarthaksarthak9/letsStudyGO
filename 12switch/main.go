package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // rand.Seed(time.Now().UnixNano()), rand.seed is now depricated
    
    // New approach:
    source := rand.NewSource(time.Now().UnixNano())
    diceNumber := rand.New(source).Intn(6)

	fmt.Println("Value of dice is", diceNumber)

	switch diceNumber{   // we are not using break
	case 1:
		fmt.Println("Move 1 step")
	case 2:
		fmt.Println("Move 2 step")
	case 3:
		fmt.Println("Move 3 step")	
	case 4:
		fmt.Println("Move 4 step")
		fallthrough					// if diceNumber is 4 it run case 5 also 
	case 5:
		fmt.Println("Move 5 step")
		fallthrough
	case 6:
		fmt.Println("Move 6 step and roll dice again")
	default:
		fmt.Println("What was that!")
			
	}

}
