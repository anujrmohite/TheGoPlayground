package main

import "fmt"

func main() {
	defer fmt.Println("Easy to Learn 01")
	defer fmt.Println("Easy to Learn 02")
	myDefer()
	fmt.Println("Defer in GoLang")
}

// it will print
// Defer in GoLang
// Easy to Learn

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)

	}
}
