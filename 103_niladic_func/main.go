// Tested in Golang Playground
//
// using an niladic function = a func that does not take in any arguments

package main

import "fmt"

func a_niladic_func() {
	fmt.Println("This is from the niladic func") // This will get printed when called
}

func main() {
	a_niladic_func()
	fmt.Println("This is from main func") // This will be printed second
}
