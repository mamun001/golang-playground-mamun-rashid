package main

import "fmt"

func main() {
	i := 8
	var pointer_to_i *int
	pointer_to_i = &i

	fmt.Println(i)
	fmt.Println(&i)
	fmt.Println(pointer_to_i)
	fmt.Println(*pointer_to_i)

}
