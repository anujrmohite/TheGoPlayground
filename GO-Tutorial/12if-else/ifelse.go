package main

import "fmt"

func main() {
	fmt.Println("if-else in golang")

	//simple and basic application
	loginCount := 23
	var res string = ""

	if loginCount < 10 {
		res = "Regular Count"
	} else if loginCount < 20 {
		res = "Avg Count"
	} else {
		res = "Non-Regular Count"
	}

	fmt.Println(res)

	//for loop can be declared like this also:

	if num := 3; num < 10{
		fmt.Println("Num is less than 10")
	}else{
		fmt.Println("Num is not less than 10")
	}
}
