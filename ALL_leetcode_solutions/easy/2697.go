// easy 2697
// lexicographically smallest palindrome
// 81.6 pct acc rate

// didn't have a clue how to solve
// got help from youtube

// JAN 4, 2025

import "fmt"


func makeSmallestPalindrome(s string) string {

ans := "" // we will append to it later
R := len(s) -1 // Right Pointer

s2 := []rune{} // because s is immutable , we will make a copy its runes


// build s2 (mutable copy)
for i := 0 ; i < len(s) ; i++ {
    s2 = append(s2, rune(s[i]))
}

//fmt.Println(s2)


// Youtube's algorithm
// compare L and R and when you see a mismatch , replace lexigraphically "bigger" one with the smaller one
for L := 0 ; L < len(s) ; L++ {
    if s2[L] > s2[R] {
        s2[L] = s2[R] // keep the small one
    } else if s2[R] > s2[L] {
        s2[R] = s2[L]
    }
    R--
} // for

//fmt.Println(s2)

// convert s2 (slice) to a string
for i := 0 ; i < len(s2) ; i++ {
    ans = ans + string(rune(s2[i]))
}


return ans
 
    
} //func
