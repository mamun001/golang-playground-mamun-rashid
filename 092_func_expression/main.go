package main

import (
	"fmt"
)

func print_foo() {
	fmt.Println("Printed The Same Thing")
}

func main() {
	f1 := print_foo
	f1()

}
