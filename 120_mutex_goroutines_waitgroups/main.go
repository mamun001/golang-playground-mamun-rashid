package main

import (
	"fmt"
	"sync"
)

// we need sync waitroups because we are using go routines

// EXAMPLE OF USING MUTEX WITH GOROUTINES

// This program will work w/o mutex , but mutex will ensure in simlar code that a common variable is not updated while a go routing is updating it

func main() {

	common_int := 0        // This is variable that will be updated by 2 goroutines
	var wg1 sync.WaitGroup // need this because we are using goroutines
	var mutex1 sync.Mutex  // this will be lock and unlock sections of code

	// 1st Goroutine
	wg1.Add(1) // adding 1 "wait"
	go func() {
		defer wg1.Done() // at the end of this func, release the waitgroup
		mutex1.Lock()    // lock updates
		common_int++     // modify the data
		mutex1.Unlock()  // unlock the lock
	}()

	// 2nd Goroutine
	wg1.Add(1)
	go func() {
		mutex1.Lock()
		defer wg1.Done()
		common_int++
		mutex1.Unlock()

	}()

	wg1.Wait() // don't finish func main until both goroutines are done
	fmt.Println(common_int)

}
