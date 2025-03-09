// easy 2695, ELO 1242
// Form Smallest Number From Two Digit Arrays

import "sort"

func is_in(ar []int, n int) bool {
    for i:=0;i<len(ar);i++ {
        if ar[i]==n {
            return true
        }
    }
    return false
}

func minNumber(nums1 []int, nums2 []int) int {

    // can we find a common digit; if more than one, then smallest
    smallest := 10 // impossible numbe in this context
    for i:=0;i<len(nums1);i++ {
        if is_in(nums2,nums1[i]) && nums1[i] < smallest {
            smallest = nums1[i]
        }
    }
    if smallest < 10 {
        return smallest
    }


    // now we know that , there is no "common" digit

    sort.Ints(nums1)
    sort.Ints(nums2)

    // now first item is the smallest digit in each
    if nums1[0] < nums2[0] {
        return nums1[0] * 10 + nums2[0]
    } else {
        return nums2[0] * 10 + nums1[0]
    }

    return 999

    
    
}
