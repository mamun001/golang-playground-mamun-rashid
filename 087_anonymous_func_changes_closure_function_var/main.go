package main

import "fmt"

func main() {

	a := 10

	closure := func() int {

		a = a + 2

		return a
	}

	fmt.Println(closure())

}
