// easy 1502 Can make arithmetic
// ELO 1154

// JAN 17, 2025

import "sort"

func canMakeArithmeticProgression(arr []int) bool {

    ans := true
    sort.Ints(arr)
    diff := arr[1] - arr[0]

    for i:=0 ; i < len(arr)-1 ; i++ {
        if arr[i+1] - arr[i] != diff {
            return false
        }
    }

    return ans
    
}

