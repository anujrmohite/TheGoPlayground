package main

import "fmt"

func main() {
	
	var username string = "string"
	fmt.Println("username")
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallFloat float64 = 822.5733435353535353
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//implicit type

	var mention_your_name = "Anuj Mohite"
	fmt.Println("\nThis is an example of implicit type: ")
	fmt.Println(mention_your_name)

	//No var style

	numberOfUsers := 300000
	fmt.Println("\nThis is no-var style type declaration")
	fmt.Println(numberOfUsers)
}
