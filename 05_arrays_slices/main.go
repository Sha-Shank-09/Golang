package main

import "fmt"

func main() {

	//Arrays
	var fruitsArr [2]string

	//Values
	fruitsArr[0] = "Jack Fruit"
	fruitsArr[1] = "Pineapple"

	//Declare + Assign
	fruitArr := [2]string{"Apple"}

	//Slices
	fruitSlice := []string{"Apple", "Jack Fruit", "Pineapple", "Cherry"}

	fmt.Println("Hello World")
	fmt.Println(fruitsArr)
	fmt.Println(fruitArr)
	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[0:2])

}
