// easy 2441 ELO 1667, Largest Positive Integer That Exists
// JAN 20, 2025


import "sort"

func contains(arr []int, n int) bool {
    ans := false
    for i:=0 ; i < len(arr) ; i++ {
        if arr[i] == n {
            return true
        }
    }

    return ans
}

func findMaxK(nums []int) int {
    
    sort.Ints(nums)
    L := len(nums)
    ans := -1

    if L == 1 {
        return -1
    }

    // start at the end with largest item and stop when have looked into all positive integers
    for i:=L-1; i > -1 && nums[i] > 0 ; i-- {
        if contains(nums,nums[i]*(-1)) { // does this number's negative exist in the array?
            ans = nums[i]
            return ans
        }
    } // for

    return ans
}

