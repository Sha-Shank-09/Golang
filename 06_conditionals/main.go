package main

import "fmt"

func main() {
	x := 5
	y := 20

	//If ELse
	if x > y {
		fmt.Printf("%d is greater than %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", y, x)
	}

	//IF ELse IF
	if x > y {
		fmt.Printf("%d is greater than %d\n", x, y)
	} else if x < y {
		fmt.Printf("%d is greater than %d\n", y, x)
	} else {
		fmt.Println("Both the numbers are the same")
	}

	//Switch
	fruit := "Mango"
	switch fruit {
	case "Mango":
		fmt.Println("Fruit is Mango")
	case "Pineapple":
		fmt.Println("Fruit is Pineapple")
	default:
		fmt.Println("No fruits available")
	}

}
