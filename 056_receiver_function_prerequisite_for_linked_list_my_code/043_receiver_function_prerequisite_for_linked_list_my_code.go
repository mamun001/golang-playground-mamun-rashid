\\ 
\\ Mamun's code
\\ tested in golang playground
\\


package main

import "fmt"

type house struct {
	number int
	street string
}

// receiver function
// THIS "h house" part coming BEFORE the rest is what makes it receiever function
// This means that this METHOD can only be applied to "house" type
//
//  This is prerequisite to doing linkeidn list the easiest way
//
func (h house) get_number() int {
	return h.number
}

func main() {
	var foo_house house
	foo_house.number = 200
	foo_house.street = "Barrington"
	n := foo_house.get_number()   // must have the last ()
	fmt.Println(foo_house)  
	fmt.Println(n)
}

