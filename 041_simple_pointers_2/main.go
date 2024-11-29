//
// POINTER LAB

package main

import "fmt"

func main() {

// declare two variables and give them values
var age int = 90
var name string = "Mamun"

fmt.Println(age)  // 90
fmt.Println(name) // Mamun

// %v = value \t = TAB %T = TYPE
fmt.Printf("%v\t%T\n", &age, &age)   // because we put & , it is pointing to the memory address; prints memory address followed by *int (* means pointer)
fmt.Printf("%v\t%T\n", &name, &name) // memory address and *string (* means pointer)
}
