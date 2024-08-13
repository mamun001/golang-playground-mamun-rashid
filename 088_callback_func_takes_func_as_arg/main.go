package main

import "fmt"

func cube(n int) int {
	return n * n * n
}

func execute_cube_func(a int, fn func(n int) int) int {

	return cube(a)

}

func main() {

	a := 10

	b := execute_cube_func(a, cube)

	fmt.Println(b)

}
