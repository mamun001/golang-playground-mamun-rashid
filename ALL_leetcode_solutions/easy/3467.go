// easy 3467, no ELO, acc rate 89%
// Transform Array by Parity

import "sort"

func transformArray(nums []int) []int {

    for i:=0;i<len(nums);i++{
        if nums[i] % 2 == 0 {
            nums[i] = 0
        }
    }

    for i:=0;i<len(nums);i++{
        if nums[i] % 2 == 1 {
            nums[i] = 1
        }
    }

    sort.Ints(nums)


    return nums
    
}
