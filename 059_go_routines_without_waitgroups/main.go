package main

import "fmt"

// func main is its own goroutine
// when you "go func" inside func main, each one of those gets its own routine
// but without waitgroups, func main can exit BEFORE those are done

// NOT using waitgroups

// If your run this, you will see that func main gets done before other go routines have a chanc
// to finish

func main() {
	fmt.Println("start of main")

	go func() {
		fmt.Println("from anonymous func 1")
	}()

	go func() {
		fmt.Println("from anonymous func 2")
	}()

	fmt.Println("end of main")

}
