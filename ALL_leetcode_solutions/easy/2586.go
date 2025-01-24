// easy 2586, ELO 1178
// Count the Number of Vowel Strings


// Jan 22, 2025

func is_vowel (r rune) bool {
    return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'  
}

func vowelStrings(words []string, left int, right int) int {

    ans := 0

    //fmt.Println(is_vowel('a'))
    //fmt.Println(is_vowel('b'))

    for i:=left ; i <= right ; i++ {
        //fmt.Println(words[i])
        a := rune(words[i][0])
        b := rune(words[i][len(words[i])-1])
        if is_vowel(a) && is_vowel(b) {
            ans++
        }
    }

    return ans
    
}

