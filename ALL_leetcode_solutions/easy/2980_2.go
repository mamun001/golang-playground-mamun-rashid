/ easy 2908, ELO 1254
// minimum sum of mountain triplets
// brute force works because len(nums)<51

func minimumSum(nums []int) int {

    min:= 151 

    for i:=0;i<len(nums);i++ {
      for j:=0;j<len(nums);j++ {
        for k:=0;k<len(nums);k++ {
            if i<j && j <k && nums[i] < nums[j] && nums[k] < nums[j] {
                sum := nums[i] + nums[j] + nums[k] 
                if sum < min {
                        min = sum
                }
            }
        }
      }
    }

    if min == 151 {
        return -1
    }

    return min
    
}
