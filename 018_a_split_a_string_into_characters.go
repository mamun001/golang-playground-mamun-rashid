//
// Mamun's code
// tested in golang playground
//
// TASK: given a string , print each charcter in string
//
package main

import (
	"fmt"
)

func main() {
	s1 := "dhalkdjald"
	for i := 0; i < len(s1); i++ {
		fmt.Println(string(s1[i]))
	}
}
