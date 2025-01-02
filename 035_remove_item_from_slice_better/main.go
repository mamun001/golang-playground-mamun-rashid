package main

import "fmt"

func main() {
slice1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
//slice2 := append(slice1[:6], 1) // works
//slice2 := append(slice1[:6], slice1[6:7]) // does not work because second arg must be an integer
//slice2 := append(slice1[:6], slice1[6:7]...) // variadic func works
slice2 := append(slice1[:6], slice1[7:]...) // [0 1 2 3 4 5 7 8 9 10 11 12], 6 removed.
fmt.Println(slice1)
fmt.Println(slice2)
}
