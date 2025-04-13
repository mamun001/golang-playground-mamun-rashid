// easy 1309, ELO 1257
// decrypt string from alphabet to integer mapping
func conv(s string) string{
    switch s {
        case "1": return "a"
        case "2": return "b"
        case "3": return "c"
        case "4": return "d"
        case "5": return "e"
        case "6": return "f"
        case "7": return "g"
        case "8": return "h"
        case "9": return "i"
        case "10#": return "j"
        case "11#": return "k"
        case "12#": return "l"
        case "13#": return "m"
        case "14#": return "n"
        case "15#": return "o"
        case "16#": return "p"
        case "17#": return "q"
        case "18#": return "r"
        case "19#": return "s"
        case "20#": return "t"
        case "21#": return "u"
        case "22#": return "v"
        case "23#": return "w"
        case "24#": return "x"
        case "25#": return "y"
        case "26#": return "z"
    }
    return ""
}

func freqAlphabets(s string) string {

    ans:=""

    // 1 = a
    // 9 = i
    // 10# = j
    // 20# = t
    // 26# = z


    // if we see 1 or 2 only then we need to see next char before deciding
    // if wee see 3-9, we can immediately say c-i

    //fmt.Println(conv("15#"))

    
    for i:=0;i<len(s);{
        //fmt.Println(i,s[i])
        if s[i] >= '3' && s[i] <= '9'{ // 3-9
            ans=ans+conv(string(s[i]))
            i++
        } else if i<len(s)-2 && s[i] == '1' && s[i+2] == '#'{ // 10#-19#
            //fmt.Println(s[i:i+3])
            ans=ans+conv(s[i:i+3])
            i=i+3
        } else if i<len(s)-2 && s[i] == '2' && s[i+2] == '#'{ // 20#-26#
            //fmt.Println(s[i:i+3])
            ans=ans+conv(s[i:i+3])
            i=i+3
        } else if i<len(s)-2 && s[i] == '1' && s[i+2] != '#'{ // 1 without #
            //fmt.Println(s[i])
            ans=ans+conv("1")
            i=i+1
        } else if i<len(s)-2 && s[i] == '2' && s[i+2] != '#'{ // 2 wtthout #
            //fmt.Println(s[i])
            ans=ans+conv("2")
            i=i+1
        } else if i>=len(s)-2 { // last 2 characters can only be 1-9
            //fmt.Println(s[i])
            ans=ans+conv(string(s[i]))
            i=i+1
        }
    }
    


    return ans
    
}
