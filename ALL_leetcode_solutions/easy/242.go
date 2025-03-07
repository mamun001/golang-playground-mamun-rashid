// easy 242
// valid anagram
// 65.8 pct acc rate

// JAN 1, 2025

import "fmt"

func isAnagram(s string, t string) bool {

    ans := true // default answer

    map1 := make(map[rune]int)
    map2 := make(map[rune]int)

    if len(s) != len(t) {
        return false
    }

    for i,_ := range s {
        map1[rune(s[i])]++
    }


    for i,_ := range t {
        map2[rune(t[i])]++
    }

    if len(map1) != len(map2) {
        return false
    }

    //fmt.Println(map1)
    //fmt.Println(map2)
    //fmt.Println(rune('a'))
    //fmt.Println(rune('e'))
    //fmt.Println(rune('z'))


    for i := 97 ; i < 122 ; i++ {
        if map1[rune(i)] != map2[rune(i)] {
            return false
        }
    }
    

    return ans
    
}

