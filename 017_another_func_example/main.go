// tested in Golang Playground

package main

import "fmt"

func square_sum(x int, y int) int {
var result int
result = x*x + y*y
return result
}

func main() {
var a int = 59
var b int = 78
fmt.Println(square_sum(a, b))
}
