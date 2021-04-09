func smallerNumbersThanCurrent(nums []int) []int {
    
    a := make([]int, len(nums))
    
    for i:=0; i < len(nums); i++ {
      for j:=0; j < len(nums); j++ {
        if nums[j] < nums[i] {
          a[i] = a[i] +1
        }
      }
    }
    
    return a
        
    
}
