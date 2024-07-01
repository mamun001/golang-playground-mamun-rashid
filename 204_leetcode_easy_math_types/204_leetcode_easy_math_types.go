
// leetcode

// output only the integer part of a number's sqrt up to 2^31 -1
// example input
// 8
// 99
// 2000000000
// output
// 2
// 9
// 44721


import (
  "fmt"
  "math"
)

func mySqrt(x int) int {
    a := float64(x)
    b := math.Sqrt(a)
    //fmt.Println(b)
    return int(b)
}
