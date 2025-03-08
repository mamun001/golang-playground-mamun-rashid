// easy 1967, ELO 1232
// Number of strings That appear as a substrings in word

import "strings"

func numOfStrings(patterns []string, word string) int {

    ans := 0

    for i:=0;i<len(patterns);i++ {
        if strings.Contains(word,patterns[i]) {
            ans++
        }
    }


    return ans
    
}
