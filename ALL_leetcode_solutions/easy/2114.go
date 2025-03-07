// easy 2114
// maximum number of words in sentences

// JAN 14, 2025


import "fmt"
import "strings"

func mostWordsFound(sentences []string) int {

    ans := 0
    count := 0 // how many words does this sentence have?

    for i:= 0 ; i < len(sentences) ; i++ {
      count = len(strings.Split(sentences[i]," ")) // how many does this one have?
      if count > ans {
        ans = count
      }
       //fmt.Println(strings.Split(sentences[0]," "))
       //fmt.Println(len(a))
    }

    return ans
    
}

