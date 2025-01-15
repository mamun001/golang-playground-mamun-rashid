// easy 1389 create target array

// JAN 15, 2025

// needed to lookup how to insert in nth place of a slice



func createTargetArray(nums []int, index []int) []int {

    ans := []int{0}


    for i:=0 ; i <len(nums); i++ {
        ans = append(ans[:index[i]+1], ans[index[i]:]...) // insert in nth place of a slice
        ans[index[i]] = nums[i]
    }
    ans = ans[:len(ans)-1] // i ended up with an extra zero in each problem
    return ans
    
}

