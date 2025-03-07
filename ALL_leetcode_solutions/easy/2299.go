// easy 2299, ELO 1242
// strong password checker II

func lowercase(s string) int {

    //fmt.Println(rune('z'))
    ans := 0

    for i:=0; i<len(s);i++ {
        if rune(s[i]) >= 97 && rune(s[i]) <= 122 {
            ans++
        }
    }

    return ans
}

func uppercase(s string) int {

    //fmt.Println(rune('A'))
    //fmt.Println(rune('Z'))
    ans := 0

    for i:=0; i<len(s);i++ {
        if rune(s[i]) >= 65 && rune(s[i]) <= 90 {
            ans++
        }
    }

    return ans
}

func digit(s string) int {

    //fmt.Println(rune('0'))
    //fmt.Println(rune('1'))
    //fmt.Println(rune('8'))
    //fmt.Println(rune('9'))
    ans := 0

    for i:=0; i<len(s);i++ {
        if rune(s[i]) >= 48 && rune(s[i]) <= 57 {
            ans++
        }
    }

    return ans
}


func special(s string) int {

    /*
    fmt.Println(rune('!'))
    fmt.Println(rune('@'))
    fmt.Println(rune('#'))
    fmt.Println(rune('$'))
    fmt.Println(rune('%'))
    fmt.Println(rune('^'))
    fmt.Println(rune('&'))
    fmt.Println(rune('*'))
    fmt.Println(rune('('))
    fmt.Println(rune(')'))
    fmt.Println(rune('-'))
    fmt.Println(rune('+'))
    */


    ans := 0

    for i:=0; i<len(s);i++ {
        if rune(s[i]) == 1 || rune(s[i]) == 33 || rune(s[i]) == 45 || rune(s[i]) == 64 || rune(s[i]) == 94 {
            ans++
        }
    }

    for i:=0; i<len(s);i++ {
        if rune(s[i]) >= 40 && rune(s[i]) <= 43 {
            ans++
        }
    }

    for i:=0; i<len(s);i++ {
        if rune(s[i]) >= 35 && rune(s[i]) <= 38 {
            ans++
        }
    }


    return ans
}

func strongPasswordCheckerII(password string) bool {

    if len(password) < 8 {
        return false
    } 

    //fmt.Println(lowercase(password))
    //fmt.Println(uppercase(password))
    //fmt.Println(digit(password))
    //fmt.Println(special(password))

    if lowercase(password) < 1 {
        //fmt.Println("lowercase")
        return false
    } 

    if uppercase(password) < 1 {
        //fmt.Println("uppercase")
        return false
    } 

    if digit(password) < 1 {
        //fmt.Println("digit")
        return false
    } 

    if special(password) < 1 {
        //fmt.Println("special")
        return false
    } 

    for i:=0; i<len(password)-1;i++ {
        if  password[i] == password[i+1] {
            fmt.Println("repeat")
            return false
        }
    }



    return true
    
}

