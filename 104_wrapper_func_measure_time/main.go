// WRAPPER FUNCTION TO SEE HOW LONG SOMETHING TOOK

package main

import (
	"fmt"
	"time"
)

func print_hello() {
	fmt.Println("hello")
}

func measure_how_long(f func()) {
	begin := time.Now()
	f()
	how_long := time.Since(begin)
	fmt.Println(how_long)
}

func main() {
	measure_how_long(print_hello)
}
