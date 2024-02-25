// we genrally use slice in golang compared to array

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome let learn slices in golang")

	var fruitList = []string{"apple", "tomato", "carrot"} // [] don't put any value inside it indicates slice , {} this represent initalization
	fmt.Printf("the type of fruitList %T\n", fruitList)

	fruitList = append(fruitList /*mention the slice here*/, "Mango", "banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1 : 3]) // 1 is inclusive where 3 is exclusive
	fmt.Println(fruitList) 


	highScores := make([]int, 4)
	
	highScores[0] = 237
	highScores[1] = 235
	highScores[2] = 236
	highScores[3] = 234

	// highScores[4] = 234 // this give error 

	highScores = append(highScores, 238, 239, 240) // it(append) reallocate the memory
	fmt.Println(highScores)
	
	sort.Ints(highScores)
	fmt.Println(highScores)



}
