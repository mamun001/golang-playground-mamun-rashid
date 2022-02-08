//
// Mamun's code
// tested in golang playground
// 
// TASK: print rune, byte, char (not a thing in golang) of a the character "d"

package main

import (
	"fmt"
)

func main() {
	char := 'd'

	fmt.Println(char)       // BYTE by default, prints 100
	fmt.Println(rune(char)) // RUNE IS ALSO A BYTE! prints 100
	fmt.Println(string(char)) // finally prints "d"

}
