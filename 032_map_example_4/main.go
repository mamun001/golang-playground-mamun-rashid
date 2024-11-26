package main

import "fmt"

func main() {

// map of string of string, basiclaly a key-value store
map1 := make(map[string]string)

// in this case, we will do name-to-role map
map1["Mamun"] = "Senior Manager"
map1["John"] = "Development Manager"
map1["Jay"] = "Product Manager"

// print the whole map
fmt.Println(map1)

// print role of Mamun
fmt.Println(map1["Mamun"]) // prints Senior Manager

// if key is not there, it prints the zero value of stringß
fmt.Println(map1["Mahfuz"]) // prints nothing

// map has a built in delete function
delete(map1, "Mamun")

// now if we try to value (role) of "Mamun", we get zero value
fmt.Println(map1["Mamun"]) // prints nothing
}
