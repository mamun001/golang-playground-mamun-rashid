// ELO 1161
// easy 2278 percentage of letter in String

// JAN 17, 2025

func percentageLetter(s string, letter byte) int {

    ans := 0

    for _,v := range s {
       if byte(v) == letter {
        ans++
       }
    }

    ans = ans * 100 / len(s)
    return ans
    
}

