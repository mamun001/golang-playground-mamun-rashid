package main

import "fmt"

// 4 parts
// 1. define the interface type
// 2. define a data structure type
// 3. implement the methods (in this case just called rent)
// 4. Use the interface in the main func

type car_rental interface {
	rent()
}

type car_model struct {
	model string
}

func (cm car_model) rent() {
	fmt.Println("Car Model Rented Is: ", cm.model)
}

func main() {
	var mycar car_rental
	mycar = car_model{"camry"}
	mycar.rent()
}
