package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	//For Loop
	i := 10
	for i > 0 {
		fmt.Println(i)
		i -= 2
	}

	//For Loop - Short
	for i := 30; i > 1; i -= 3 {
		fmt.Println(i)
	}

	//FizzBuzz
	for i := 1; i <= 30; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
