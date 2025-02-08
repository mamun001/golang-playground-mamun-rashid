// 2255 ELO 1261
// easy
// Count Prefixes of a Given string

func is_prefix(p string, s string) bool {
    ans := true
    np := len(p)
    ns := len(s)

    if ns < np {
        return false
    }

    for i:=0;i<np;i++ {
        if p[i] != s[i] {
            //fmt.Println("false at ",i)
            return false
        }
    }

    return ans
}

func countPrefixes(words []string, s string) int {

    ans := 0

    //fmt.Println(is_prefix("abc","abcdefgjjjjjj"))

    for i,_ := range words {
        if is_prefix(words[i],s) {
            ans++
        }
    }


    return ans
    
}

