package main

import "fmt"

//func add_two(a int, b int) int {
//return a + b
//}

func main() {
	a := 10
	b := 20

	c := func(a int, b int) int {
		return a + b
	}(a, b)

	fmt.Println(c)

}
