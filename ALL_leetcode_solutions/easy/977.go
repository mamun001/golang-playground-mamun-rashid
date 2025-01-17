// easy 977
// ELO 1129

// JAN 16, 2025

import "sort"

func sortedSquares(nums []int) []int {

    for i:= 0 ; i < len(nums) ; i++ {
        nums[i] = nums[i] * nums[i]
    }

    sort.Ints(nums)
    
    return nums
}
