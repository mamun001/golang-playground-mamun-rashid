
// Mamun's code
// Tested in go playground
// 
// TASK: write a function does a to the 4, when is a is the inout to the function

package main

import "fmt"

func tothefour(x int) int {
	return x * x * x * x
}

func main() {
	var a int = 5
	fmt.Println(tothefour(a))
}



