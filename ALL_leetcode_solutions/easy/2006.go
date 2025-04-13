// easy 2006, ELO 1272
// count number of pairs with absolute difference K

func abs(n int) int {
    if n < 0 {
        return n * -1
    }
    return n
}


func countKDifference(nums []int, k int) int {

    ans:=0

    for i:=0;i<len(nums)-1;i++ {
        for j:=i;j<len(nums);j++ {
            if abs(nums[i] - nums[j]) == k {
                ans++
            }
        }
    }
    return ans
    
}
