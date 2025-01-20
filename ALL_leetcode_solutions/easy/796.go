// easy 796 rotate string
// ELO 1167
// JAN 18, 2025

// spent 30-45 on 1 algorithm, that failed more many cases, i.e. did not work
// then , I was forced to come up with plan B , which turned out much better
// I am legend! You can't keep me down!

import "strings"


func rotateString(s string, goal string) bool {

    ans := false

    padded_string := ""

    if len(s) != len(goal) {
        return false
    }

    for i,_ := range s {
        padded_string = s[i:] + s[0:i]
        if padded_string == goal {
            return true
        }
    }
    
    return ans
    
}
