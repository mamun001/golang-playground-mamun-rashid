// easy 747, ELO 1988
// Largest Number At Leat

// Jan 24, 2025

import "sort"

func dominantIndex(nums []int) int {
    max := 0
    max_index := -1
    second := 0
    L := len(nums)
    ans := -1

    for i:=0 ; i<len(nums); i++ {
        if nums[i] > max {
            max = nums[i]
            max_index = i
        }
    }
    fmt.Println(max,max_index)
    sort.Ints(nums)

    second = nums[L-2]
    fmt.Println("second",second)

    if max < (second * 2) {
        return -1 
    } else {
        return max_index
    }
    
    return ans

}
