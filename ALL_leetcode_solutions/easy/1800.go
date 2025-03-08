// easy 1800 , ELO 1229
// maximum ascending subarray sum

func maxAscendingSum(nums []int) int {

    current_sum := nums[0]
    max_sum := nums[0]
    

    for i:=1;i<len(nums);i++ {

        if nums[i] > nums[i-1] {
            current_sum = current_sum + nums[i]
            if current_sum > max_sum {
                max_sum = current_sum
            }
        } else {
            current_sum = nums[i]
        }
    }


    return max_sum
    
}
