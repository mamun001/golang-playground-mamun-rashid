// easy 1935 , ELO 1227
// Maximum Number of Words You Can Type

import "strings"

func canBeTypedWords(text string, brokenLetters string) int {

    words := strings.Split(text," ")

    map1 :=make(map[rune]int)

    for _,v := range brokenLetters {
        map1[rune(v)] = 1
    }

    //fmt.Println(words)
    //fmt.Println(map1)

    
    bad_count := 0
    for i,_ := range words {
        for _,x := range words[i] {
            if map1[rune(x)] == 1 {
               //fmt.Println(words[i])
               bad_count++
               break
            }
        }
    }

    ans := len(words) - bad_count
    

    return ans
    
}

