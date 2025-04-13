// easy 2697
// lexicographically smallest palindrome
// 81.6 pct acc rate
// code from 3rd try (first w/o help)



// JAN 4, 2025, first attempt with youtube help
// March 20, again needed help from google search
// March 24, done w/o help, but took 50 minutes

func is_palin(s string) bool {
    l:=0
    r:=len(s)-1

    for l<r {
        if s[l] != s[r] {
            return false
        } 
        l++
        r--
    }
    return true
}


func makeSmallestPalindrome(s string) string {

    //fmt.Println(is_palin("abccdba"))

    if is_palin(s) {
        return s
    }

    l:=0
    r:=len(s)-1 

        if s[l] == s[r] {
            new_string := string(s[l])+makeSmallestPalindrome(s[l+1:len(s)-1])+string(s[r])
            return(new_string)
            //fmt.Println(new_string)
        }
        if s[l] < s[r] {
            new_string := s[:len(s)-1]+string(s[l])
            return(makeSmallestPalindrome(new_string))
            //fmt.Println(new_string)
        }
        if s[l] > s[r] {
            new_string := string(s[r]) + s[1:len(s)]
            return(makeSmallestPalindrome(new_string))
            //fmt.Println(new_string)
        }
    
    


    return ""


    
} //func
