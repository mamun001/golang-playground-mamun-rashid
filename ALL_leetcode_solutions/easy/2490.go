// easy 2490, ELO 1263
// circular sentence

import "strings"

func isCircularSentence(sentence string) bool {

    words:=strings.Split(sentence," ")
    //fmt.Println(words)

    for i:=0;i<len(words)-1;i++ {
        //fmt.Println(len(words[i]))
        L:=len(words[i])
        //fmt.Println(words[i][L-1])
        if words[i][L-1] != words[i+1][0] {
            return false
        }
    }

    last_word:=words[len(words)-1]
    //fmt.Println(last_word)
    LL:=len(last_word) // length of the last world
    //fmt.Println(last_word[LL-1])

    if last_word[LL-1] != words[0][0] {
        return false
    }

    return true
    
}
