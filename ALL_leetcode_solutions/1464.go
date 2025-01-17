// easy 1464 maximum product
// ELO 1121

// JAN 16, 2025

import "sort"
import "fmt"

func maxProduct(nums []int) int {

    ans := 0
    sort.Ints(nums)

    L := len(nums)

    biggest := nums[L-1]
    fmt.Println(biggest)
    second_biggest := nums[L-2]
    fmt.Println(second_biggest)

    ans = (biggest-1) * (second_biggest-1)

    return ans
    
}

