package main

import "fmt"

// EXAMPLE OF INTERFACE TYPE ASSERTION

func main() {

	var empty_interface_1 interface{} = "hello again"

	// This will panic because the value that the empty interface holds is a string, not an integer
	// int1 := empty_interface_1.(int)
	// The following will work, because value that the empty interface holds is a string
	// s1 := empty_interface_1.(string)

	// What if you did no know what type it was ??
	// That's where type assertion comes into play
	// golang allows you to use "ok" idiom to test the type w/o panic-ing (crashing)

	var_int, ok := empty_interface_1.(int)
	if ok {
		fmt.Println("type is int::", var_int)
	}

	var_string, ok := empty_interface_1.(string)
	if ok {
		fmt.Println("type is string::", var_string)
	}

}
