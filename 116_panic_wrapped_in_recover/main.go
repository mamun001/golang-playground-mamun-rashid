package main

import "fmt"

// SIMPLE PANIC WITH RECOVER

// for readibility, making the panic func a separate func
func cause_a_panic() {
	// panic must have at least 1 arg which is the "message"
	panic("something bad happened. panic-ing")
}

func recover_from_a_panic() {
	//
	// recover must be part of a defer , otherwise it will never be called, because panic would have already crashed the program
	// NOTE: func that panics (or "may" panic in real life, must be warpped inside the recovery func, which makes sense if you think about it)
	//       so, in this case recover_from_a_panic func is wrapping the cause_a_panic func

	// The way you know that this whole thing worked is that , when you run this, you will not see RAW panic message,
	//   instead you will see a nice message saying a panic happened and it will nicely report what the panic-ing func sent as a dying message
	//

	defer func() {
		recover_output := recover() // we capture what the recover send us, which will be nil if there is no panic
		// you can see that if you comment out the cause_a_panic func call
		if recover_output == nil {
			fmt.Println("There was no panic detected") // you will see this if you comment out the cause_a_panic call below
		} else {
			// recover_output is not nil, that means there was a panic
			fmt.Println("recovered from panic: panic-ed code sent the following message::", recover_output)
		}
	}()
	cause_a_panic()
}

func main() {
	// call the func that wraps a panic-ing func
	recover_from_a_panic()
}
