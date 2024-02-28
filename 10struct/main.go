package main

import "fmt"

func main() {
	fmt.Println("Struct in GoLang")
	sarthak := User{"Sarthak", "sarthak@go.dev", true, 20}
	fmt.Println(sarthak)
	fmt.Printf("This way give you more details: %+v\n", sarthak)
	fmt.Printf("Name is %v, email is %v\n", sarthak.Name, sarthak.Email)

}

type User struct {  // here User first letter need to be in Capital it will export out or access by anyone
	Name   string   // fields first letter need to be Capital, so that it access by anyone
	Email  string
	Status bool
	Age    int
}
