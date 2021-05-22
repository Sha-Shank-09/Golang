package main

import (
	"fmt"
	"strconv"
)

//Define car struct
type Car struct {
	// brandName string
	// carName   string
	// year      int
	// mileage   int

	brandName, carName string
	year, mileage      int
}

// Value Reciever Method
func (c Car) honk(lane int) string {
	return "Please move aside for " + c.carName + " on " + strconv.Itoa(lane)
}

// Pointer Reciever Method
func (c *Car) addMileage(distance int) {
	c.mileage = c.mileage + distance
}

func main() {

	car1 := Car{brandName: "VW", carName: "Polo", year: 2014, mileage: 21422}
	car2 := Car{"Mazda", "CX-5", 2021, 1290}

	fmt.Println(car1)
	fmt.Println(car2)
	fmt.Println(car1.brandName == car2.brandName)
	fmt.Println(car1.honk(5))
	fmt.Println("The mileage on Car2 is " + strconv.Itoa(car2.mileage))
	car2.addMileage(40)
	fmt.Println("The mileage on Car2 is " + strconv.Itoa(car2.mileage))
}
