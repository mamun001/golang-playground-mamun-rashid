// easy 1417, ELO 1242
// Reformat the string


func abs(n int) int {
    if n < 0 {
        return n * -1
    }
    return n
}


func is_digit(b byte) bool {
    if b > 47 && b < 58 {
        return true
    }
    return false   
}

func is_letter(b byte) bool {
    if b > 96 && b < 123 {
        return true
    }
    return false   
}



func is_ok(s string) bool {
    
    for i:=0;i<len(s)-1;i++ {
        if is_digit(s[i]) && is_digit(s[i+1]) {
            return false
        }
        if is_letter(s[i]) && is_letter(s[i+1]) {
            return false
        }
    }
    return true
}

func reformat(s string) string {

    digits:=0
    letters:=0
    only_digits:=""
    only_letters:=""

    for i:=0;i<len(s);i++ {
        //fmt.Println(s[i])
        if s[i] > 47 && s[i] < 58 {
            digits++
            only_digits = only_digits + string(s[i])
        } else {
            letters++
            only_letters = only_letters + string(s[i])
        }
    }

    //fmt.Println(digits,letters)

    // mathematically true
    if abs(digits-letters) > 1 {
        return ""
    }


    // if s is already formatted correctly
    if is_ok(s) {
        return s
    }

    //fmt.Println(only_digits)
    //fmt.Println(only_letters)

    
    ans:=""
    if len(only_digits) == len(only_letters) {
        for i:=0;i<len(only_digits);i++ {
            ans = ans + string(only_digits[i])
            ans = ans + string(only_letters[i])
        }
    }


    // if there are 1 extra digit than letters
    if len(only_digits) > len(only_letters) { 
        ans = ans + string(only_digits[len(only_digits)-1])
        for i:=0;i<len(only_letters);i++ {
            ans = ans + string(only_letters[i])
            ans = ans + string(only_digits[i])
        }
    }

    // if there are 1 extra letter than digits
    if len(only_digits) < len(only_letters) { 
        ans = ans + string(only_letters[len(only_letters)-1])
        for i:=0;i<len(only_digits);i++ {      
            ans = ans + string(only_digits[i])
            ans = ans + string(only_letters[i])
        }
    }


    return ans
    
}
