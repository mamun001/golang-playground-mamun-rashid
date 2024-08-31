package main

// MULTIPLE defers (LIFO)

//
import (
	"fmt"
)

// When you run this, you will see that prints happens in reverse order
func main() {

	defer fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
	defer fmt.Println("d")

}
