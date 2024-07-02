package main

import (
	"fmt"
)

func main() {
	channel1 := make(chan string)

	go func() {
		channel1 <- "This message is going to channel1"
	}()

	retrieved_message := <-channel1

	fmt.Println(retrieved_message)

}
