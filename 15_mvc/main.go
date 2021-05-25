package main

import (
	"fmt"

	"github.com/shashank-kakarla/go_basics/15_mvc/models"
)

func main() {
	user := models.User{
		ID:        2,
		FirstName: "ShankS",
		LastName:  "Akagami",
	}
	fmt.Println(user)
}
