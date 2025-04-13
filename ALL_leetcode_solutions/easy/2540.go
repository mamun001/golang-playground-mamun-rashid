// easy 2540, ELO 1250
// Minimum Common Value
// nums1/nums2 can have 10,000 values
// so, efficieny matters

func getCommon(nums1 []int, nums2 []int) int {

    i:=0
    j:=0 

    for i<len(nums1) && j<len(nums2) {
        if nums1[i] == nums2[j] {
            return nums1[i]
        }
        if nums1[i] < nums2[j] {
           i++
        } else if nums1[i] > nums2[j] {
           j++
        }
    }

    return -1
    
}
