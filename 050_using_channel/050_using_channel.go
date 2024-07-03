package main

import (
	"fmt"
)

func main() {

	channel1 := make(chan string)

	go func() {
		channel1 <- "This message will go into channel 1"
	}()

	message_out := <-channel1

	fmt.Println(message_out)

}
