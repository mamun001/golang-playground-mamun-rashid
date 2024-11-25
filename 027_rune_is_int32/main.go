package main

import "fmt"

// RUNE IS AN ALIAS for int32!
var r1 rune = 'a'
var r2 rune = 97

func main() {

fmt.Println(r1)         // prints 97
fmt.Println(r2)         // also prints 97
fmt.Println(string(r1)) // prints a
fmt.Println(string(r2)) // also prints a
}
