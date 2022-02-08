//
// Mamun's code 2.7.2022
// tested in golnag playground
//
// TASK: implement a stack of integers (create, print, push, pop)
//



package main

import "fmt"

func create_stack() []int {
	result := make([]int, 0)
	return result
}

func push_to_stack(stack []int, new_item int) []int {
	// in my implementation, HIGHER END on SLICE is th TOP of the Stack
	//
	// for example [100,200,300]. = 100 was pushed first, then 200, then 300
	//
	result := stack
	result = append(result, new_item)
	return result
}

func pop_stack(stack []int) []int {
	result := stack[:len(stack)-1] // make a new slice with len of 1 less than the stack fed to this function
	return result
	// NEED TO ADD HERE A CONDITION WHERE STACK IS ALREADY EMPTY
}

func main() {

	foo_stack := create_stack()
	fmt.Println(foo_stack)                    // prints []
	foo_stack = push_to_stack(foo_stack, 100) // prints [100]
	fmt.Println(foo_stack)
	foo_stack = push_to_stack(foo_stack, 200) // prints [100,200]
	fmt.Println(foo_stack)
	foo_stack = push_to_stack(foo_stack, 300) // prints [100,200,300]
	fmt.Println(foo_stack)

	foo_stack = pop_stack(foo_stack)
	fmt.Println(foo_stack) // prints [100,200]

}

