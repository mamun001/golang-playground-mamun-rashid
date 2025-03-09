// easy 2016, ELO 1246
// Maximum Difference Between ...
func maximumDifference(nums []int) int {

    
    max:=0

    for i:=0;i<len(nums)-1;i++ {
        for j:=i+1;j<len(nums);j++ {
          if nums[j] - nums[i] > max {
            max = nums[j] - nums[i]
          }
        }
    }

    if max == 0 {
        return -1
    }
    return max
    
}
