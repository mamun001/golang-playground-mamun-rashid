// easy 2974 , ELO 1184
// minimum number game

// Jan 25, 2025

import "sort"


func numberGame(nums []int) []int {

    ans :=[]int{}

    sort.Ints(nums)

    for i:=0 ; i < len(nums) ; i=i+2 {
        //fmt.Println(nums[i+1],nums[i])
        ans = append(ans,nums[i+1])
        ans = append(ans,nums[i])
    }


    return ans
    
}

