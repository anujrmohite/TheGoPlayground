package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("switcha and case in golang")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of Dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Value : 1")
		break
	case 2:
		fmt.Println("Value : 2")
		break
	case 3:
		fmt.Println("Value : 3")
		break
	case 4:
		fmt.Println("Value : 4")
		break
	case 5:
		fmt.Println("Value : 5")
		break
	case 6:
		fmt.Println("Value : 6")
		break
	default:
		fmt.Println(nil)
		break

		//##########################
		// we can also add fallthrough
		//##########################

	}
}
