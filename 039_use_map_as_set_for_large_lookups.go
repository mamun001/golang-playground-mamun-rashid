package main

import "fmt"

func main() {

// making a map of string who value is struct of nothing is teh best way to implement a set-like lookup
// this is very useful for leetcode problems where performance of essence AND you have to do lookup from very large set
cities := make(map[string]struct{})

cities["nyc"] = struct{}{}     // INSERT
cities["la"] = struct{}{}      // INSERT
cities["chicago"] = struct{}{} // INSERT

fmt.Println(cities) // map[chicago:{} la:{} nyc:{}]

delete(cities, "la") // delete is built-in func for maps
fmt.Println(cities) // map[chicago:{} nyc:{}]

_, answer := cities["dallas"] // When you access a value from a map, Go provides an optional second return value, which is a boolean indicating whether the key exists in the map
fmt.Println(answer) // false

_, answer = cities["nyc"]
fmt.Println(answer) // true
}

