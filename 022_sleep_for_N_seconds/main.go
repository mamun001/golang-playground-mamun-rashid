package main

// We need  time package to use sleep
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	// really complicated syntax
	time.Sleep(5 * time.Second)
	fmt.Println(time.Now())

}
