

// modified from gobyexample.com


package main

import "fmt"

func main() {

    // create a new map whose KEY is string and value is INT
    m := make(map[string]int)


    // insert some values in the map
    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)


    // create an integer from a key-value from map m
    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    fmt.Println("len:", len(m))


    // delete one key-value pair from m
    delete(m, "k2")
    fmt.Println("map:", m)


    // _  is a blank identifier  (to see if the key k2 exists or not. we deleted k2 above.
    _, prs := m["k2"]
    fmt.Println("prs:", prs)


    // create new map on the fly
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}


