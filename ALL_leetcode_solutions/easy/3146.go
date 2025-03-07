// 3146 easy Permutation Difference between Two strings

// acc rate 87.9%

// DEC 24, 2024

import "fmt"


func abs (i int) int {
    ans := 0
    if i < 0 {
        ans = i * -1
    } else {
        ans = i
    }

    return ans
}
func index_of_letter(s string, l rune) int{

// finds the index of a letter inside a string
    ans := 0
    for i,v := range s {
        if v == l {
            ans = i
        }
    }
    return ans
}



func findPermutationDifference(s string, t string) int {
    
    //fmt.Println(index_of_letter("majun", 'j'))

    ans := 0

    for i,v := range s {
        ans = ans + abs (i - index_of_letter(t,v))
    }

    return ans
}
