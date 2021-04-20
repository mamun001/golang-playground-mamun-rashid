package main

// Mamun Rashid
// Tested in go Playground

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
