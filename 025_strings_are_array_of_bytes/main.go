package main

import "fmt"

// simple string
var s string = "Hello Mamun"

// slice of bytes ; byte = 8 bits ; 0-255; empty right now
var byte_slice []byte

func main() {

// giving it 3 elements . 0-255 are possible values
byte_slice = []byte{65, 1, 35}

// convert string S into slice of bytes ; prints the "elements" of that slice
fmt.Println([]byte(s)) // [72 101 108 108 111 32 77 97 109 117 110]

// prints as is [65 1 35]
fmt.Println(byte_slice) // [65 1 35]
}
