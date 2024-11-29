

// TESTED IN GOLANG PLAYGROUND

package main

import "fmt"

// simplest possible example of bit shift operator <<
// 1 in binary is 00000001
// when we use << operator , it shifts the "1" one bit to the left, so it becomes 00000010
// that equals 2 in decimal

func main() {
	fmt.Println(1 << 1) // prints 2
}
