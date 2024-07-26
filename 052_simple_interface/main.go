package main

import "fmt"

// write a interface type
type rental_car interface {
	drive()
}

// write a struct type
type car struct {
	make string
}

// write code to implement the interface
func (c car) drive() {
	fmt.Println("Driving a", c.make)
}

func main() {
	// create a value of the struct type
	var myCar rental_car = car{"Toyota"}
	// use the value of the struct type to call the method
	myCar.drive()
}
