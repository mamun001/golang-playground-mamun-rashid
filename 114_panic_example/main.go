package main

// SIMPLE PANIC WITHOUT RECOVER

//

func main() {

	// panic must have at least 1 arg which is the "message"
	panic("something bad happened, calling panic")
}
