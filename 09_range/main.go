package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	//Slices or Arrays
	ids := []int{1, 2, 3, 4, 5, 6}

	//Looping with index
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//Looping without index
	//Wihtout the index variable
	for id := range ids {
		fmt.Printf("ID: %d\n", id)
	}
	//Using _
	for _, id := range ids[0:2] {
		fmt.Printf("ID: %d\n", id)
	}

	//Maps
	emails := map[string]string{"ShankS": "shanks@email.com", "Monkey": "monkey@email.com"}

	//Looping with K,V
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
	//Looping with only K
	for k := range emails {
		fmt.Printf("Key is %s\n", k)
	}
	//Looping with only V
	for _, v := range emails {
		fmt.Printf("Value is %s\n", v)
	}

}
