package main

import "fmt"

func main() {
	fmt.Println("GO functions")
	greater()

	greater2()
	res := adder(23,23)
	fmt.Println(res) 
}

func greater() {
	fmt.Println("Namaste")
}

func greater2() {
	fmt.Println("Anther Namaste")
}

func adder(num1 int ,num2 int ) int {
	fmt.Println("The return Value is: ")
	return num1 + num2 
}