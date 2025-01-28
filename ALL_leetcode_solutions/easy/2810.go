// easy 2810, ELO 1192
// faulty keyboard

// JAN 25, 2025

func rev(s string) string {
    ans := ""
    n := len(s)
    for i:=n-1; i>-1;i-- {
        ans = ans + string(s[i])
    }

    return ans
}

func finalString(s string) string {

    ans := ""

    //fmt.Println(rev("abcdefghji"))

    n := len(s)

    for i:=0; i<n; i++ {
        if s[i] != 'i' {
            ans = ans + string(s[i])
        } else { // char is i
            ans = rev(ans)
        }
    }

    return ans
    
}

