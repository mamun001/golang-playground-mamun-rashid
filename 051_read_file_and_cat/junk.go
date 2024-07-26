package main

import "fmt"

# define an interaface
type rental_car interface {
	drive()
}

# define a struct for the interface above
type car struct {
	name string
}

# implement the interface
func (c car) drive() {	
	fmt.Println(c.name, "is driving")
}	

func main() {
	var my_car rental_car
	my_car = car{"Toyota"}
	my_car.drive()
}	
	


