package main

import "fmt"

func main()  {
	fmt.Println("Let's Learn pointers");

	// var ptr *int
	// fmt.Println("Value of Pointer is", ptr);

	myNumber := 23;
	var ptr = &myNumber;

	fmt.Println("Value of actula pointer is", ptr);
	fmt.Println("Value of actula pointer is", *ptr);

	*ptr = *ptr + 2;

	fmt.Println("New value of actula pointer is", myNumber);

}