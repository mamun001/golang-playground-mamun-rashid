

// Passed December 17 2024

func twoSum(nums []int, target int) []int {
    s := make([]int, 2)
    for i :=0 ; i < len(nums); i++ {
        for j :=0; j < len(nums); j++ {
         if i == j {
           break
         } else if nums[i] + nums[j] == target{
             s[0] = i
             s[1] = j
         }
      }
    }
    return s
}
