// easy 2956, Find Common Elements Between
// ELO 1215

func is_in(ar []int, n int) bool {

    for i:=0;i<len(ar);i++ {
        if ar[i] == n {
            return true
        }
    }

    return false
}

func findIntersectionValues(nums1 []int, nums2 []int) []int {

    ans1:= 0
    ans2:= 0

    for i:=0;i<len(nums1);i++ {
        if is_in(nums2, nums1[i]) {
            ans1++
        }
    }

    for i:=0;i<len(nums2);i++ {
        if is_in(nums1, nums2[i]) {
            ans2++
        }
    }

    ans := []int{ans1,ans2}

    return ans
    
}
