// An user input program that checks for various user errors

package main

// we need the following packages
// bufio = to read from terminal
// fmt = to write to terminal
// os = stdin
// strconv = to convert string to integer
// strings = to trim space from  terminal input

// THIS HAD TO TESTED on local, because golang playground would exit without a termnal input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	read_from_terminal := bufio.NewReader(os.Stdin) // open a new buffer from terminal
	fmt.Print("Enter an Interger: ")
	termnal_input, _ := read_from_terminal.ReadString('\n') // read the string from terminal
	termnal_input = strings.TrimSpace(termnal_input)        // get rid of any spaces

	int_from_terminal, err := strconv.Atoi(termnal_input)   // convert to int
	if err != nil {
		fmt.Println("You did not enter a valid integer. Please enter a valid integer next time")
	return
		}

	fmt.Println(int_from_terminal)
}


