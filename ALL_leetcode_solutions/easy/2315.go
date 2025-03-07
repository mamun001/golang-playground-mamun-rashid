// easy 2315, ELO 1251
// Count Asterisks

func countAsterisks(s string) int {

    in_pipes := false
    ans := 0

    for i:=0 ; i<len(s);i++ {
        if s[i] == '|' && in_pipes == false {
            in_pipes = true
            
        } else if s[i] == '|' && in_pipes == true {
            in_pipes = false         
        }

        if s[i] == '*' && in_pipes == false {
            ans++
        }


    }



    return ans
    
}

