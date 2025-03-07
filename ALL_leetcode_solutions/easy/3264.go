// easy 3264
// Final Array State After K multiplication operations I

// acc rate 87.6
// JAN 13, 2025

import "fmt"

func index_of_min (arr []int) int {

    min := arr[0]+1 // initial value
    ans := 0

    for i := 0 ; i < len(arr); i++ {
        if arr[i] < min {
            min = arr[i]
            ans = i // arr[i] MUST BE less than current minimum for the index to change and so, if there are repeats, 1st one is kept
        }

    }
    return ans  
}


func getFinalState(nums []int, k int, multiplier int) []int {

indx := 101 // initial crazy value


 for i:=1 ; i <= k ; i++ { // do this k times
    //fmt.Println(index_of_min(nums))
    // find the index of the min and save
    indx = index_of_min(nums)
    fmt.Println(indx)
    // multiply that item with the multip
    nums[indx] = nums[indx] * multiplier
    //fmt.Println(nums)
 }
return nums   
}

