
// Test in Golang Playground
// Program to use different of variables and print their value and types
package main

import "fmt"

var a int = 8
var b float32 = 12.5
var c bool = false
var d string = "Mamun"

// %t = tab , %T = type , %v = value , \n = new line
func main() {
fmt.Printf("%v\t%T\n", a, a) // 8, int
fmt.Printf("%v\t%T\n", b, b) // 12.5 float32
fmt.Printf("%v\t%T\n", c, c) // false bool
fmt.Printf("%v\t%T\n", d, d) // Mamun string
}
