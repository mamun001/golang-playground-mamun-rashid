//  Mamuns code
// tested in Go playground
//
// TASK : make a key value pair using golnag's map and add value
//
package main

import "fmt"

func main() {

	map2 := make(map[string]string) // Basically key value pair. index(key) is string, value is also a string
	map2["mamun"] = "Barrington Street"   // key = "mamun" value = "Barrington Street"
	fmt.Println(map2["mamun"])
}
