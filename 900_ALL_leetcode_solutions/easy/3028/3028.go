// easy 3028 
// ant on the boundary

// ELO 1115 (4th lowest)

// took 5 minute to write and submit

func returnToBoundaryCount(nums []int) int {

    ans := 0
    sum := 0

    for i:= 0 ; i < len(nums); i++ {
        sum = sum + nums[i]
        if sum == 0 {
            ans++
        }
    }
    
    return ans
}
