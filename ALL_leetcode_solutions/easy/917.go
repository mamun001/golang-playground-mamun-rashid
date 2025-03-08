// easy 917, ELO 1229
// Needed help from youtube

func is_english(r rune) bool {
    if r >=65 && r<=90 || r>=97 && r<=122 {
        return true
    }
    return false
}

func reverseOnlyLetters(s string) string {

    // A = 65
    // Z = 90
    // a = 97
    // z = 122

    //fmt.Println(byte('A'))

    ar :=[]rune(s)
    //fmt.Println(ar)

    l:=0
    r:=len(ar)-1

    for l<r {
        //fmt.Println(is_english(ar[l]))
        //fmt.Println(is_english(ar[r]))
        if is_english(ar[l]) == true && is_english(ar[r]) == true{
            ar[l],ar[r] = ar[r],ar[l]
            l++
            r--
        } else if is_english(ar[l]) == true && is_english(ar[r]) == false {
            r--

        } else if is_english(ar[l]) == false && is_english(ar[r]) == true {
            l++
        } else if is_english(ar[l]) == false && is_english(ar[r]) == false {
            l++
            r--
        }
    }

    s = string(ar)
    //fmt.Println(s)



    return s
    
}
