package main

import (
	"fmt"
)

func main() {
	fmt.Println("Map in Golang")

	language := make(map[string]string)

	language["JS"] = "javascript"
	language["RB"] = "ruby"
	language["py"] = "python"

	fmt.Println("Here is the language that I learned: ", language)
	fmt.Println("JS full form: ",language["JS"])

	delete(language, "RB") // we can use same for slice
	fmt.Println("Here is the language that I learned: ", language)

	// Loops are interseting in goLang

	for key, value := range language{   // to ignore any (key, value use _) like for _, value := range language{}
		fmt.Printf("For key %v, the value is %v\n", key, value)
	}

}
