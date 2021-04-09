func numIdenticalPairs(nums []int) int {
    var good int = 0
    
    for i:=0; i < len(nums)-1; i++ {
        for j:=i+1; j < len(nums); j++ {
            if nums[i] == nums[j] {
              good = good + 1
            }
        }
    }
    
    return good
    
}
