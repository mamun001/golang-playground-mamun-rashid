// 68% acc rate

// easy, 1455 Check if a word occurs as a prefix of any word in a sentence

// DEC 26, 2024

import "fmt"
import "strings"


func isPrefixOfWord(sentence string, searchWord string) int {

    ans := -1
    
    words := strings.Split(sentence," ") // split into array of strings
    //fmt.Println(temp)

    for i,v := range words {
        if strings.HasPrefix(v,searchWord) {
            //fmt.Println("hello")
            ans = i+1
            break
        }
    }

    return ans
}


