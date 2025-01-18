// import "fmt"

func isAcronym(words []string, s string) bool {

    ans := true

    L1 := len(words)
    L2 := len(s)
    //fmt.Println(len(words),len(s))

    if L1 != L2 { // they have to exact same length at the minimum
        return false
    } else {
        for i :=0 ; i<L1 ; i++ {
            if words[i][0] != s[i] { // no match ? exit
                ans = false
            }
        }
    }

    return ans
    
}

