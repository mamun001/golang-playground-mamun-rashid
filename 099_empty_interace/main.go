package main

import "fmt"

// EMPTY INTERFACE

func main() {

	var i interface{}

	i = 5
	fmt.Println(i)

	i = "hello"
	fmt.Println(i)

	i = 5.5
	fmt.Println(i)

	i = true
	fmt.Println(i)

}
