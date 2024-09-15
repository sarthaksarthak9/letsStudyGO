package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study in golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // if you want to know present date time day you have to go with this way only

	createDate := time.Date(2023, time.January, 12, 21, 27, 0, 0, time.UTC)
	
	fmt.Println(createDate)

	

}
