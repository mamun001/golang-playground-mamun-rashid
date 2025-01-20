// easy 3162 Find the number of good pairs
// ELO 1168.7

// JAN 18,2025

func numberOfPairs(nums1 []int, nums2 []int, k int) int {

    ans := 0

    for i:=0 ; i<len(nums1); i++ {
        for j:=0 ; j<len(nums2); j++ {
            if nums1[i] % (nums2[j] * k) == 0 {
                ans++
            }
        }
    }

    return ans
    
}
