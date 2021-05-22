package main

import (
	"fmt"
)

func main() {

	//Define a Map<Key, Value>
	contact := make(map[string]int)

	//Assign Values
	contact["ShankS"] = 101010101
	contact["Monkey"] = 9090999909

	fmt.Println(contact)
	fmt.Println(len(contact))
	fmt.Println(contact["Monkey"])

	//Delete Value from Map
	contact["WB"] = 12312345
	fmt.Println(contact)
	delete(contact, "WB")

	//Declare and Assign K,V
	emails := map[string]string{
		"ShankS": "shanks@email.com",
		"Monkey": "monkey@email.com"}
	emails["WB"] = "WB@email.com"
	fmt.Println(emails)

}
