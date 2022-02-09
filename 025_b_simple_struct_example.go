//
// tested in golang playground
//
// Mamun's code
//
// TASK: define a struct and create one of that tye and add values and print
//

package main

import "fmt"

type house struct {
	number int
	street string
	city   string
	state  string
	zip    int
}

func main() {
	var house_foo = house{200, "Barrington Way", "Santa Clarita", "CA", 91350}
	fmt.Println(house_foo)
}

