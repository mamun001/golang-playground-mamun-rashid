package main

import "fmt"

func main() {

	a := 10

	closure := func() int {
		return a + 2
	}

	fmt.Println(closure())

}
