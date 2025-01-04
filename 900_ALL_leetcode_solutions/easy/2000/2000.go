// easy 2000
// reverse prefix of word
// 86.3 acc rate
// JAN 3, 2025


func reversePrefix(word string, ch byte) string {

    ans := ""
    found_at := -1

    for i,v := range word {
        if v == rune(ch) {
           found_at = i
           break
        } 
    }

    if found_at == -1 {
        return word
    }

    for i := found_at ; i > -1 ; i-- {
        ans = ans + string(word[i])
    }

    for i := found_at+1 ; i < len(word) ; i++ {
        ans = ans + string(word[i])
    }

    return ans
    
}

