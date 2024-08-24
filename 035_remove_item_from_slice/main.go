package main

import "fmt"

func main() {

	var index int = 2

	s1 := make([]int, 0)

	s1 = append(s1, 10)
	s1 = append(s1, 20)
	s1 = append(s1, 30)
	s1 = append(s1, 40)
	s1 = append(s1, 50)

	fmt.Println(s1)

	s1 = append(s1[0:index], s1[index+1:]...)

	fmt.Println(s1)

}
