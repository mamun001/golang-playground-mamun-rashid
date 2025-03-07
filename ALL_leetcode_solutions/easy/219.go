// easy, 219, ELO unknown, acc rate 47%!
// finally got one without help!
// Contains Duplicate II

func containsNearbyDuplicate(nums []int, k int) bool {

    ans := false
    for i:=0;i<len(nums)-1;i++ {
        for j:=i+1;j<=i+k && j<len(nums);j++{
            if nums[i] == nums[j] {
                return true
            }
        }
    }

    return ans
    
}
