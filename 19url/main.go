package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&payment=ghj45jhj"

func main() {
	fmt.Println("Welcome to handling urls in golang")
	fmt.Println(myurl)

	// parsing (parsing - extract some info)

	result, _ := url.Parse(myurl) // url is lib
	fmt.Println(result.Scheme)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Host)

	qparams := result.Query();

	fmt.Printf("The type of qparams: %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams{
		fmt.Println("The params: ",val)
	} 

	partsOfURL := &url.URL{ // pass a reference thats y &
		Scheme: "https",
		Host: "lco.dev",
		Path: "tutcss" ,
		RawPath: "user=Sarthak",
	}

	anotherUrl := partsOfURL.String()

	fmt.Println(anotherUrl)
}
