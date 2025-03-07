// easy 1876, ELO 1249
// substrings of size three with 

func countGoodSubstrings(s string) int {

    if len(s) < 3 {
        return 0
    }


    ans := 0

    for i:=0;i<len(s)-2;i++ {
        rep := 0
        if s[i] == s[i+1] {
            rep++
        }
        if s[i+1] == s[i+2] {
            rep++
        }
        if s[i] == s[i+2] {
            rep++
        }
        if rep == 0 {
           ans++
        }



    }

    return ans
    
}
