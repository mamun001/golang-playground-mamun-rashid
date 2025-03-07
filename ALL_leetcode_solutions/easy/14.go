// #14 , easy
// 44.4 pct acc rate

// DEC 28, 2024



//import "fmt"


func prefix_of_two(s1 string, s2 string) string {
    var runes []rune // we build the common prefix of two string by comparing 1 rune (char) at a time
    //runes = append(runes,'a')
    //runes = append(runes,'b')

    loop_length := 0 // loop only runs N times where N is the length of the smaller of the two strings

    if len(s1) <= len(s2) {
        loop_length = len(s1)
    } else {
        loop_length = len(s2)
    }

    for i := 0 ; i < loop_length ; i++ {
        if s1[i] == s2[i] {
            runes = append(runes,rune(s1[i])) // building the prefix once rune at a time
        } else {
            break
        }
    }

    return string(runes)
}

func longestCommonPrefix(strs []string) string {

    ans := strs[0]
    //ans := prefix_of_two(strs[0],strs[1])
    //s3 := prefix_of_two ("ab", "ac") // testing
    //fmt.Println(s3)
    

    // we figure out the common prefix of the 1st and the 2nd word
    // then we take that prefix and make a new one with 3rd word
    // and we continue till the end of the words
    for i := 0 ; i < len(strs)-1 ; i++ {
        ans = prefix_of_two(ans,strs[i+1])
        //fmt.Println(i,strs[i],strs[i+1],ans)
    }

    return ans
    
} // END OF FUNC longestCommonPrefix
