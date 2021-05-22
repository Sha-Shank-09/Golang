package main

import "fmt"

func main() {

	//Everything in GO is pass-by-value
	//Pointers
	a := 1
	b := &a
	fmt.Println(a, b)
	fmt.Printf("%T --- %T\n", a, b)

	// Use * to get the value from the address (a == *b = True)
	fmt.Println(a)
	fmt.Println(*b)

	// Update the value of variable with the pointer
	*b = 5
	fmt.Println(a)
}
