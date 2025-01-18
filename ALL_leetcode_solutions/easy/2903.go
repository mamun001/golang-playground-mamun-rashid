// ELO 1157
// easy 2903 find indices with

// Jan 17, 2025


func abs(n int) int {
    ans := n
    if n < 0 {
        ans = n * -1
    }
    return ans
}

func findIndices(nums []int, indexDifference int, valueDifference int) []int {

    ans :=[]int{}

    for i:=0 ; i < len(nums); i++ {
        for j:=0 ; j <len(nums); j++ {
          if abs(i - j) >= indexDifference && abs(nums[i] - nums[j]) >= valueDifference {
            ans = append(ans,i)
            ans = append(ans,j)
          }

        }
    }

    if len(ans) == 0 {
        ans = append(ans,-1)
        ans = append(ans,-1)
    }
    
    return ans
}

