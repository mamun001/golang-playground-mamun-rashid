package main

import (
	"fmt"
	"time"
)

func print_something(s string) {
	fmt.Println(s)
}

func main() {
	go print_something("Message 1")
	go print_something("Message 2")
	go print_something("Message 3")
	go print_something("Message 4")
	go print_something("Message 5")
	go print_something("Message 6")
	go print_something("Message 7")

	time.Sleep(2 * time.Second)
}
