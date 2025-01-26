// easy 1122, ELO 1188
// relative sort array

// Jan 24, 2025

import "sort"

func is_in(ar []int, n int) bool {
    ans := false
    for i:=0;i<len(ar); i++ {
        if ar[i] == n {
            return true
        }
    }
    return ans
}

func relativeSortArray(arr1 []int, arr2 []int) []int {

    ans := []int{}
    spares := []int{}

    // fill ans with all the ones found in arr2
    for i:=0; i<len(arr2); i++ {
        for j:=0;j<len(arr1); j++ {
            if arr1[j] == arr2[i] {
               ans = append(ans,arr1[j])
            }
        }
    }

    // find arr1-arr2 "extra ones"
    for i:=0; i<len(arr1); i++ {
       if is_in(arr2,arr1[i]) != true {
          spares = append(spares,arr1[i])
       }
    }

    //fmt.Println(spares)
    sort.Ints(spares)

    // pad ans with the "extra one sorted"
    for i:=0; i<len(spares); i++ {
       ans = append(ans,spares[i])
    }

    return ans
    
}
