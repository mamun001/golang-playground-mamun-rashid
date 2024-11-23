package main

import "fmt"

func XtoY(x int, y int) int {
	var result int = x

	for i := 1; i < y; i++ {
		result = result * x
	}
	return int(result)

}

func main() {
	var x int = 7
	var y int = 4
	var result = 1
	result = XtoY(x, y)

	fmt.Printf("%v\n", result)

}
