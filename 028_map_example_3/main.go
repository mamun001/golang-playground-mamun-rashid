package main

import "fmt"

func aleady_there(s1 []string, first_name string) bool {
	result := false

	for _, v := range s1 {
		if v == first_name {
			result = true
			fmt.Println(s1)
			//fmt.Println(first_name)
			break
		}
	}

	return result
}

func is_there_a_repeat_first_name(data map[string]string) bool {
	slice1 := make([]string, 0)
	result := false

	for first := range data {
		if aleady_there(slice1, data[first]) == true {
			result = true
		} else {
		}
		slice1 = append(slice1, data[first])

	}
	return result
}

func main() {
	first_last := make(map[string]string)
	first_last["john"] = "d"
	first_last["paul"] = "e"
	first_last["peter"] = "f"
	first_last["jose"] = "g"
	first_last["kumar"] = "h"

	result := is_there_a_repeat_first_name(first_last)

	fmt.Println(result)

}
