package main

// Mamun Rashid's code
// Tested in go Playground

// Write a function that takes in integer and retruns true if it is 1, otherwise returns false
//

import (
	"fmt"
)

func return_if_1(myint int) bool {
	if myint == 1 {
		return true
	} else {
		return false
	}
}

func main() {

	returned_bool := return_if_1(3)
	fmt.Println(returned_bool)

}
