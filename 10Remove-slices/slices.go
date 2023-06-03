package main

import (
	"fmt"
//	"sort"
)

func main() {
//	fmt.Println("Welcome to slices")
	var fruitlist = []string{"Apple", "Tomato", "Peach"}
//	fmt.Println("Type of fruitlist is %T\n", fruitlist)

	//how to add the values
	fruitlist = append(fruitlist, "Mango" , "Banana")

	//Printing values after adding it to the list 

//	fmt.Println("Type of fruitlist is %T\n", fruitlist)

	// var highscore = [10]int{1,2,3,4,5,6,67,7,9};

	highscore1 := make([]int, 10);
	
	highscore1[0] = 0
	highscore1[1] = 1
	highscore1[2] = 2
	highscore1[3] = 3
	highscore1[4] = 4
	highscore1[5] = 5
	highscore1[6] = 6
	highscore1[7] = 7
	highscore1[8] = 8
	highscore1[9] = 9

//	fmt.Println("The list is", highscore1);

	highscore1 =append(highscore1, 11,12,13,14,15,161,7,171111118)

	// fmt.Println("The list is", highscore1);

	// fmt.Println(sort.IntsAreSorted(highscore1));

	fmt.Println("#############")
	fmt.Println("how to remove slices from slices based on index")

	var courses = []string{
		"reactjs",
		"js",
		"swift",
		"ruby",
		"python",
	};

	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index],courses[index+1:]...)

	fmt.Println(courses)

	// fruitlist = append(fruitlist[1:])
	// fmt.Println(fruitlist)

	// fruitlist = append(fruitlist[1:])
	// fmt.Println(fruitlist)

	// fruitlist = append(fruitlist[1:])
	// fmt.Println(fruitlist)

	// fruitlist = append(fruitlist[1:])
	// fmt.Println(fruitlist)

	// fruitlist = append(fruitlist[1:])
	// fmt.Println(fruitlist)



}