// Mmauns code
// tested in go playground
//
// TASK: append to an empty slice 
//
package main

import "fmt"

func main() {

	s2 := make([]int, 0)
	s2 = append(s2, 100, 200)
	fmt.Println(len(s2))
	fmt.Println(s2[0])
	fmt.Println(s2[1])
}
