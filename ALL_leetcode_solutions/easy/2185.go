// easy 2185 ELO 1167 Counting Words with

// Jan 20,2025

import "strings"

func prefixCount(words []string, pref string) int {

    ans := 0

    for _,v := range words {

        a := strings.Index(v,pref)
        if a == 0 {
            ans++
        }
        //fmt.Println(a)

    }

    return ans
    
}
