package main

import "fmt"

func main(){
	fmt.Println("If-Else in goLang")

	loginCount := 23

	if loginCount < 20{
		fmt.Println("Push a little harder")
	}else if loginCount < 40{
		fmt.Println("Good Job")
	}else{
		fmt.Println("Need to work")
	} 

	// This we do a lot while handling web req

	if num := 3; num < 10{  // most of the we get value from browser and check it on the go
		fmt.Println("Number is less than 10")
	}else{
		fmt.Println("Number is greater than 10")
	}

	// To check a value is nil

	// if err := nil{
		
	// }
}





