package main

// difference between methods and funct is that, we use and create function.
// when we define function into classes and struct that we call them methods

import "fmt"

func main() {
	fmt.Println("Struct in GoLang")
	sarthak := User{"Sarthak", "sarthak@go.dev", true, 20}
	fmt.Println(sarthak)
	fmt.Printf("This way give you more details: %+v\n", sarthak)
	fmt.Printf("Name is %v, email is %v\n", sarthak.Name, sarthak.Email)
	sarthak.getStatus()
}



type User struct {  // here User first letter need to be in Capital it will export out or access by anyone
	Name   string   // fields first letter need to be Capital, so that it access by anyone
	Email  string
	Status bool
	Age    int
}

func(u User) getStatus(){
	fmt.Printf("The user status is %v \n", u.Status)
}



