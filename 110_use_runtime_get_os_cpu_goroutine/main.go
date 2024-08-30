package main

// We need  runtime package to get these info
import (
	"fmt"

	"runtime"
)

func main() {
	fmt.Println("Arch", runtime.GOARCH)
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())

}
