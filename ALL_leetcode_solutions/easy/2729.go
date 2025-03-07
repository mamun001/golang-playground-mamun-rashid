// easy 2729, ELO 1228
// check if The Number is Fascinating

import "strconv"
import "strings"


func is_fasc(s string) bool {
    ans:= true
    if strings.Contains (s, "0") {
       return false
    }
    if strings.Count (s, "1")  != 1 {
       return false
    }
    if strings.Count (s, "2")  != 1 {
       return false
    }
    if strings.Count (s, "3")  != 1 {
       return false
    }
    if strings.Count (s, "4")  != 1 {
       return false
    }
    if strings.Count (s, "5")  != 1 {
       return false
    }
    if strings.Count (s, "6")  != 1 {
       return false
    }
    if strings.Count (s, "7")  != 1 {
       return false
    }
    if strings.Count (s, "8")  != 1 {
       return false
    }
    if strings.Count (s, "9")  != 1 {
       return false
    }


    return ans

}

func isFascinating(n int) bool {

    //fmt.Println(strconv.Itoa(n))
    //fmt.Println(strconv.Itoa(n*2))
    //fmt.Println(strconv.Itoa(n*3))

    s := strconv.Itoa(n) + strconv.Itoa(n*2) + strconv.Itoa(n*3)
    //fmt.Println(s)

    //fmt.Println(is_fasc("11"))


    return is_fasc(s)
    
}
