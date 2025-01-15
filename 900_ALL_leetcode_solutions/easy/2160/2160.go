// easy 2160 minimum sum of four digit ....

// JAN 14, 2025

// my ALG: beast 100% cpu and 70% RAM!

import "fmt"
import "sort"

func minimumSum(num int) int {

    ans := 0

    // easy to get the digits , since num is always a 4 digit number
    d1 := num / 1000
    d2 := (num - d1*1000 ) / 100
    d3 := (num - d1*1000 - d2*100) / 10
    d4 := num - d1*1000 - d2*100 - d3*10
    fmt.Println(d1,d2,d3,d4)

    // put the digits into a slice so we can sort them
    slice1 := []int{d1,d2,d3,d4}
    sort.Ints(slice1)
    fmt.Println(slice1)

    // MY ALG: answer is always the two smallest digits * 10 plus the other two digits
    ans = slice1[0]*10+slice1[1]*10+slice1[2]+slice1[3]

    return ans
    
}

