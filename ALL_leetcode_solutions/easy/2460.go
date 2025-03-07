/*
easy, 2460, ELO 1224
Apply Operations to an array

*/

func applyOperations(nums []int) []int {

    // do the operations
    for i:=0;i<len(nums)-1;i++{
        if nums[i] == nums[i+1] {
            nums[i] = nums[i] * 2
            nums[i+1] = 0
        }
    }

    
    // if we see a zero, cut it out of the array and keep count
    zeros:=0
    for i:=0;i<len(nums);{
        if nums[i] == 0  {
            nums = append(nums[:i],nums[i+1:]...)
            zeros++
        } else {
          i++
        }
    }

    // append the zeros at the end
    for i:=0;i<zeros;i++{
        nums = append(nums,0)
    }
    

    return nums
    
}
