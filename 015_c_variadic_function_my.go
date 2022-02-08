// Mamun's code
// tested in go lang playground
// Click here and start typing.
package main

import "fmt"

func variadic_mult(sent_ints ...int) int {
	result := 1
	for i := range sent_ints {
		result = result * sent_ints[i]
	}

	return result
}

func main() {
	fmt.Println(variadic_mult(5, 10, 20))
}

