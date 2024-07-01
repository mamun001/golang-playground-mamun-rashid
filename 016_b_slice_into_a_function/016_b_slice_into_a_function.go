//
// Mamun's code
// tested on golang playground
//
// TASK : feed a function a slice and return the sum of all its elements
//
package main

import "fmt"

func sum_of_slice(s []int) int {
	result := 0
	for i := 0; i < len(s); i++ {
		result = result + s[i]
	}
	return result
}

func main() {
	nums := make([]int, 3)
	nums[0] = 5
	nums[1] = 20
	nums[2] = 100
	fmt.Println(sum_of_slice(nums))
}

