// easy 2418, ELO 1193
// sort the people

// Jan 26, 2025

// my alg poor: 11% + 23%

import "sort"

// get the index of the given value

func idx (ar []int, n int) int {
    ans := -1

    for i:=0; i<len(ar); i++ {
        if ar[i] == n {
            return i
        }
    }

    return ans
}

func sortPeople(names []string, heights []int) []string {

    ans :=[]string{}
    copy :=[]int{}

    n := len(names)

    // make a copy so we can sort
    for i:=0; i<n; i++ {
        copy = append(copy,heights[i])

    }

    // sort the copy
    sort.Ints(copy)

    // largest to smallest, find the index in the original
    index := 0
    for i:=n-1;i>-1;i--{
        //fmt.Println(idx(heights,copy[i]))
        index = idx(heights,copy[i])
        ans = append(ans,names[index])
    }

    return ans
    
}
