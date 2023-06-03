package main

import "fmt"

func main() {
	fmt.Println("Structurs in golang")
	// There in no inheritence in golang

	anuj := User{
		"Anuj",
		"anuj@go.dev",
		true,
		23,
	}
	fmt.Println("Anuj Mohite's Details are: ")
	fmt.Println("Name is", anuj.NickName)

}

//Here uppercase has to be included

type User struct {
	NickName   string
	Email  string
	Status bool
	Age    int
}
