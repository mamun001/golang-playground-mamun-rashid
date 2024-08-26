// ANONYMOUS STRUCT

package main

import "fmt"

func main() {
	magnus := struct {
		name   string
		rating int
	}{
		name:   "magnus",
		rating: 2800,
	}

	fmt.Println(magnus)
}
