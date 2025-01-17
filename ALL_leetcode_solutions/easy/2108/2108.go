// easy 2108 
// 83.9 acc rate

// Find first palindromic string is the array

// JAN 4, 2025

import "fmt"

func is_palin(s string) bool {
    ans := true

    R := len(s) -1
    for _,v := range s {
        if v != rune(s[R]) {
            return false
        }
        R--
    }

    return ans
}

func firstPalindrome(words []string) string {
    
  ans := ""

  //fmt.Println(is_palin("abcdefghhgfedcba"))

  for _,v := range words {
    if is_palin(v) {
        return v
    }
  }

  return ans

}
