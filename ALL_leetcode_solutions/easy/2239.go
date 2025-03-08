// easy 2239, ELO 1241
// find closest number to zero

func abs(n int) int {
    if n < 0 {
        return n* (-1)
    }
    return n
}


func findClosestNumber(nums []int) int {

    n:= len(nums)
    i:=0
    ans:=100001 // based on max possible value

    for i<n {
        if abs(nums[i]) < abs(ans) {
            ans = nums[i]
        } else if abs(nums[i]) == abs(ans) && nums[i] > 0 {
            ans = nums[i]
        }
        //fmt.Println(i,ans)
        i++
    }


    return ans
    
}
