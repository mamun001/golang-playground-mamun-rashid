// Tested in Golang Playground
//
// using an init function in a package

package main

import "fmt"

func init() {
	var a int = 10
	fmt.Println("This is from init func") // This will print first
	fmt.Println(a)
}

func main() {
	// fmt.Println(a)  commented out because it will error out , because a is not define. Does not carry over from init func
	fmt.Println("This is from main func") // This will be printed second
}
