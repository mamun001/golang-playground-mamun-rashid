// Leetcode Easy # 1704 Determine if String Halves Are Alike

// Strings

/* My data set
"abcdefgh"
"aeiosh"
"mamunm"
"rashid"
"abcabc"
"aeiouZaeiouY"
"godfatheri"
"alpacino"
"scentofawomanP"
"BooKii"
*/

func count_vowels(s string) int {
    answer := 0
    for _, char := range s {
        if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'{
            answer ++
        } 
        if char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U'{
            answer ++
        }
    } // for
    
    return answer
}



func halvesAreAlike(s string) bool {
    
    answer := true
    
    half := len(s)/2
    //fmt.Println(half)
    
    s1:=s[0:half]
    s2:=s[half:len(s)]
    
    if (count_vowels(s1)) == (count_vowels(s2)) {
        answer = true 
    } else {
        answer = false
    }
    
    return answer
    
}
