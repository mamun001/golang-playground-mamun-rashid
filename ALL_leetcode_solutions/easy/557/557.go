// 557
// easy 
// Reverse Words in a String III

// JAN 4, 2024

import "fmt"
import "strings"

func reverse_string (s string) string {

    ans := ""

    for i := len(s) -1 ; i > -1 ; i-- {
        ans = ans + s[i:i+1]
    }

    return ans
}


func reverseWords(s string) string {

  ans := ""

  //fmt.Println(reverse_string("A'BMTRYYYYab"))

  split_string := strings.Fields(s) // split into array of strings using space as delimeter
  //fmt.Println(split_string)
  
  for i,v := range split_string {
    if i == len(split_string)-1 { // dont add space if it is the last word
       ans = ans + reverse_string(v)
    } else {
        ans = ans + reverse_string(v) + " "
    }
  }

  return ans
}

