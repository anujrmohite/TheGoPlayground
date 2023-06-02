package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in Go-lang")

	var fruitlist [4]string

	fruitlist[0] = "apple"
	fruitlist[1] = "tomato"
	fruitlist[2] = "peach"
	fruitlist[3] = "mango"

	fmt.Println("Fruit list is", fruitlist)
	fmt.Println("and length is", len(fruitlist))

	var veglist [3]string;
	veglist = [3]string{"Potato", "beans", "mushroom"}
	fmt.Println("Veglist is following:")
	fmt.Println(veglist)
	fmt.Println("and length of veglist is:")
	fmt.Println(len(veglist))
}

