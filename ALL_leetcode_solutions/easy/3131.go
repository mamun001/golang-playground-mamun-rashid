// easy 3131 find the integer added
// ELO 1160

// JAN 17, 2025

import "sort"

func addedInteger(nums1 []int, nums2 []int) int {

    sort.Ints(nums1)
    sort.Ints(nums2)

    ans := nums2[0] - nums1[0]

    return ans

    
}
