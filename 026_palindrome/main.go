// IS PALINDOME??

package main

import "fmt"

func main() {

	// This code has been tested in Golang Playground

	// input data for testing
	var input_string string = "Mamun"
	//var input_string string = "MaM"
	//var input_string string = "M"
	//var input_string string = "Mam"
	//var input_string string = "Rashid"

	var is_palindrome bool = true // starts with true and then gets false if there is one mismatch

	// 0 to 4 (if length of input_string is 5
	for i := 0; i < len(input_string); i++ {
		//fmt.Println(i)
		if input_string[i] != input_string[len(input_string)-i-1] {
			is_palindrome = false // already a mismatch, exit the lopp
			break
		}

		// fmt.Println(input_string[i])                     // 0 TO 5-1 (4)
		// fmt.Println(input_string[len(input_string)-i-1]) // 5-0-1 (4) TO 5-4-1 (0)
	}

	if is_palindrome {
		fmt.Println(input_string, "  ", "True")
	} else {
		fmt.Println(input_string, "  ", "False")
	}
}
