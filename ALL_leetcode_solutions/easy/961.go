// ELO 1161
// easy 961 N-Repeated element

// JAN 17, 2025

import "sort"

func repeatedNTimes(nums []int) int {

    // because there are n+1 uniq elements, that means ONLY 1 element appears more than once
    // that's the one to find

    // we sort, tha will put those next to each other
    // that all we have to check is if the current one is the same as the last one

    ans := 10001 // impossible value to start with 

    sort.Ints(nums)
    

    for i:=1 ; i < len(nums); i++ {
        if nums[i] == nums[i-1] {
            ans = nums[i]
        }
    }

    return ans
    
}
