package main

import (
	"fmt"
)

type car = string

func drive_a_car(c1 car) string {
	result := "I am driving " + c1
	return (result)
}

func main() {

	car1 := "Toyota Celica"

	fmt.Println(drive_a_car(car1))

}
