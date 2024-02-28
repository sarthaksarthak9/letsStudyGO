package main

import "fmt"

func main(){
	fmt.Println("Loops in Golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thrusday", "Friday", "Saturday"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++{ // we don't have ++i
	// 	fmt.Println(days[i])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	for index, day := range days{
		fmt.Printf("index %v day %v\n", index, day)
	}

	rougheValue := 1;

	for rougheValue < 10{                    // similiar to while loop
		if rougheValue == 8{
			break;
		}

		if rougheValue == 4{               // it skip value 6
			rougheValue++
			continue  
		}

		if rougheValue == 9{
			goto lco
		}

		fmt.Println("Value is : ", rougheValue)
		rougheValue++;
	}

	lco: // define label for goto
		fmt.Println("Jumping at learningcoding.in")
}