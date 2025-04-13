// easy 2475, ELO 1256
// number of unequal triplets in array
// brute force works because array length 100 or less

func unequalTriplets(nums []int) int {

    ans:=0
    for i:=0;i<len(nums);i++ {
      for j:=0;j<len(nums);j++ {
        for k:=0;k<len(nums);k++ {
            if i < j && j < k && nums[i] != nums[j] && nums[i] != nums[k] && nums[j] != nums[k] {
                ans++
            }
        }
     }
    }

    return ans
    
}
