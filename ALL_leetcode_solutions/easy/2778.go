// easy 2778 sum of squares
// ELO 1151

// JAN 17, 2025

func sumOfSquares(nums []int) int {

    ans := 0
    n := len(nums)

    for i:=0 ; i < len(nums); i++ {
        if n % (i+1) == 0 { // 1 indxed-array
            ans = ans + nums[i] * nums[i]
        }
    }
    
    return ans
}
