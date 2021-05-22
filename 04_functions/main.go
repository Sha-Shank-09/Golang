package main

import (
	"fmt"
	"math"
)

func greeting(name string) string {
	return "Hello" + name
}

func sum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(math.Ceil(2.7))
	fmt.Println(sum(3, 4))
	fmt.Println("Shashank")
}
