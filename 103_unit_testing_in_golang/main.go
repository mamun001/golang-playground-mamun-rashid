// TESTING IN GO

package main

import "fmt"

func cube(a int) int {
	return a * a * a
}

func main() {
	fmt.Println(cube(20))
}
