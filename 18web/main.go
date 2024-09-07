package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com/"

func main() {
	fmt.Println("Google web server")
	
	response, err := http.Get(url)

	if err != nil{
		panic(err)
	}

	fmt.Printf("The type of the response is %T\n", response)
	defer response.Body.Close()  // caller's responsibility to close connection

	databytes, err := ioutil.ReadAll(response.Body) // ioutil majorly use for reading

	if err != nil{
		panic(err)
	}

	content := string(databytes)
	fmt.Println(content)
}
