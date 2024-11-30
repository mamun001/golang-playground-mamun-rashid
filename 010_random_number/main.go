
//
// Tested on Golang Playground

// will generate an integer between 0 and 999
package main

import (
"fmt"
"math/rand"
)

func main() {

var r int = 0

r = rand.Intn(1000) // will print between 0 and 999

fmt.Println(r)

}
