package main

import (
	"fmt"
)

func main() {

	string1 := "Abcdefg"
	fmt.Println(string1)

	string2 := string(string1[0])
	fmt.Println(string2)

	rune1 := []rune(string1)
	fmt.Println(rune1)

	fmt.Printf("%c\n", rune1[0])

}
