// easy 88 , acc rate 52%
// not sure why it works even though I did not modify nums1, a pointer thing??

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int)  {

    ar := append(nums1[:m],nums2...)

    sort.Ints(ar)
    
    
}
