// easy 1394 find lucky integer
// 5th or so easiest ELO

// Jan 16, 2025

import "fmt"

func findLucky(arr []int) int {

    ans := 0

    freq := []int {}

    n := len(arr)

    // lazy programming
    for i:=0 ; i<501 ; i++ {
        freq = append(freq,0)
    }

    //fmt.Println(arr)
    //fmt.Println(freq)

    for i:=0 ; i<n ; i++ {
        //fmt.Println(i)
        freq[arr[i]]++
    }
    //fmt.Println(freq)

    for i:=0 ; i<501 ; i++ {
        if freq[i] == i {
           ans = i
           //fmt.Println(i)
        }
    }

    if ans == 0 {
        ans = -1
    }

    return ans
   
}
