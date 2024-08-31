package main

// MESSAGE PASSING BETWEEN GOROUTINES

// Need sync for workgroups , so main func does not finish before goroutines do.
import (
	"fmt"
	"sync"
)

func main() {

	// Need workgroups (see above)
	var wg sync.WaitGroup
	wg.Add(2)

	// create the channel
	message_channel := make(chan string)

	// Goroutine 1 sends the message
	go func() {
		fmt.Println("goriutine 1 says: ", "Added a message")
		message_channel <- "Super Secret Message"
		defer wg.Done()
	}()

	// Goroutine 2 picks up the message
	go func() {
		m1 := <-message_channel
		fmt.Println("goriutine 2 says: ", m1)
		defer wg.Done()
	}()

	wg.Wait()

}
