package main

import "fmt"

// 122: FUNC RETURING A FUNC

func func_that_returns_a_func() func() string {
	return func() string {
		return "Chess"
	}

}

func main() {

	func1 := func_that_returns_a_func()

	fmt.Println(func1())

}
