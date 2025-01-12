// easy 2089 Find Target Indices After Sorting Array
// acc rate 77%
// I knew the algorithm (BST), BUT HAD TO LOOKUP GOLANG CODE in YOUTUBE

// my code : cpu beats 100%, RAM beats 9%

import "fmt"
import "sort"


func targetIndices(nums []int, target int) []int {

    ans := []int{}
    sort.Ints(nums)

    //fmt.Println(nums)

    L:=0
    R:=len(nums)-1
    pos := 0 // what position are we checking right now?
    pointer := 0
    
    for L <= R { // I didn't know you could write loop like this
       pos = (L+R) / 2
       //fmt.Println(L,pos,R)
       if nums[pos] == target {
        ans = append(ans,pos)
        break
       } else if nums[pos] > target { // look to the left
         R = pos - 1
       } else if nums[pos] < target { // look to the right
         L = pos + 1
       }
    } // for loop end

    // pos now has at least one target
    // BUT PROBLEM REQUIRES THAT WE RETURN ALL INDEXES


    // look for repeats to the left
    pointer = pos
    for nums[pointer] == target && pointer != 0 {
        //fmt.Println("looking left")
        if nums[pointer-1] == target {
            ans = append(ans,pointer-1)
        }
        pointer--
    }

    // look for repeats to the right
    pointer = pos
    for nums[pointer] == target && pointer != len(nums)-1 {
        //fmt.Println("looking right")
        if nums[pointer+1] == target {
            ans = append(ans,pointer+1)
        }
        pointer++
    }

    sort.Ints(ans) // problem expects sorted positions




    return ans
    
}

