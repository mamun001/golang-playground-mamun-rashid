package main

import "fmt"

func get_index_of_value(s []int, value int) int {
	var result int = -1
	for i := range s {
		if s[i] == value {
			result = i
		}
	}

	return result
}

func main() {

	slice1 := make([]int, 0)

	slice1 = append(slice1, 10)
	slice1 = append(slice1, 20)
	slice1 = append(slice1, 30)
	slice1 = append(slice1, 40)
	slice1 = append(slice1, 50)

	fmt.Println(slice1)

	fmt.Println(get_index_of_value(slice1, 5))

}
