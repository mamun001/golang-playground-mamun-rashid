// easy 2733 niether minimum nor maximum
// ELO 1130-ish
// JAN 16, 2025


import "sort"

func findNonMinOrMax(nums []int) int {

    ans := 0
    sort.Ints(nums) // sorting makes code super simple
    L := len(nums)

    if L == 1 || L == 2{
        return -1 // no such number can exist
    } else {
        return nums[1] // if L > 2 , always return the 2nd element
    }
   
    return ans
    
}

