// easy 217
// acc rate 62.6

// Jan 1, 2025

// My submission falied because it took too long
// I had to look up the right algorithm on youtube

import "sort"



func containsDuplicate(nums []int) bool {
    //s1 := []int {}
    sort.Ints(nums)
    ans := false // default ans

    for i := 0 ; i < len(nums) -1 ; i++ {
        //fmt.Println(v, does_slice_contain(s1,v))
        if nums[i] == nums[i+1] {// next neighbor is of same value
           return true
        }
    }

    //fmt.Println(s1)
    return ans
    
}

