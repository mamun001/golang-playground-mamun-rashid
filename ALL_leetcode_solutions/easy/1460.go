// easy 1460 
// Make two arrays equal
// ELO 1151

// JAN 17, 2025

// fansy mumbo jumbdo , but essentially two arays have to same items


import "sort"

func canBeEqual(target []int, arr []int) bool {

    sort.Ints(target)
    sort.Ints(arr)

    L := len(arr)

    ans := true

    for i:=0 ; i < L ; i++ {
       if arr[i] != target[i] {
        ans = false
       }
    }

    return ans
    
}

