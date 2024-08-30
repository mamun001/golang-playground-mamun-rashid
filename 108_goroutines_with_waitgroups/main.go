package main

// We need sync pack to use Waitgroup
import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("start of main")

	// We will define A wiatgroup just like you would define an int
	var wg sync.WaitGroup

	// We will tell wg that that it has 2 items in it
	wg.Add(2)

	go func() {
		// we don't want "done" to be executed until this block is totall done
		defer wg.Done()
		fmt.Println("from anonymous func 1")
	}()

	go func() {
		// we don't want "done" to be executed until this block is totall done
		defer wg.Done()
		fmt.Println("from anonymous func 2")
	}()

	// Without the explicit "wait" func main will get finished even if wg is not done
	wg.Wait()
	fmt.Println("end of main")

}
