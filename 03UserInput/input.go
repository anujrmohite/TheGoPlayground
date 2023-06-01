package main

import (
	"bufio"
	"fmt"
	"os"
)

//Remembar almost everything is type in golang
//Now lets explore bufio
func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza")

	//comma ok syntax - comma error syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("The type of this rating is %T \n", input)

}
