package main

import "fmt"

func main() {
	fmt.Println("Go loops")
	days := []string{
		"Sunday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	fmt.Println(days)

	// lOOP 1:

	for d := 0; d < len(days); d++ {
		fmt.Println("Position : ", d, days[d])
	}

	//lOOP 2:

	for i := range days {
		fmt.Println(days[i])
	}

	rougeValue := 1
	for rougeValue < 10 {
		fmt.Println("value is: ", rougeValue)
		rougeValue++
	}
}
