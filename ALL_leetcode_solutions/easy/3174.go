// easy 3174 ELO 1255
// Clear digits


func is_digit(r rune) bool {
    if r == '0' || r == '1' || r == '2' || r == '3' || r == '4' || r == '5' || r == '6' || r == '7' || r == '8' || r == '9' {
        return true
    } else {
        return false
    }

    return false
}

func clearDigits(s string) string {

    ans := ""

    //fmt.Println(is_digit(rune(ans[0])))

    for i,val := range s {
        //fmt.Println(is_digit(rune(val)))
        if is_digit(rune(val)) {
            ans = ans[:len(ans)-1] // pop last
        } else {
            ans = ans + string(s[i]) // add in this char
        }
    }


    return ans
    
}
