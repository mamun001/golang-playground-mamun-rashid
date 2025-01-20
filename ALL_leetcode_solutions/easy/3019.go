// easy 3019
// number of changing keys
// ELO 1175

// JAN 19, 2025

import "fmt"

func abs (n int) int {
    ans := n
    if n < 0 {
        ans = ans * -1
    }
    return ans
}

func countKeyChanges(s string) int {

    ans := 0

    //fmt.Println(rune('A')) // 65
    //fmt.Println(rune('a')) // 97 , 32 more
    //fmt.Println(rune('Z')) // 90
    //fmt.Println(rune('z')) // 122, 32 more

    for i:=1 ; i<len(s) ; i++ {
        diff := abs(int(s[i])-int(s[i-1]))
        fmt.Println(diff)
       if diff != 0 && diff != 32 {
        ans++
       }
    }
    

    return ans
    
}
