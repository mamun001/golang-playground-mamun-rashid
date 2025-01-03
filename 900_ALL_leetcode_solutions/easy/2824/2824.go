// easy 2824
// acc rate
// 2 pointers

// JAN 3 , 2025



func countPairs(nums []int, target int) int {

    ans := 0

    for L := 0 ; L < len(nums)-1 ; L++ { 
        for R := L+1 ; R < len(nums) ; R++ {
            if nums[L] + nums[R] < target {
               ans++ 
            }
        }
    }

    return ans
    
}
