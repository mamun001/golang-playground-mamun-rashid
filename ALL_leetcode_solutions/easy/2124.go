// easy 2124 ELO 1203 
// check if all A's appear

func checkString(s string) bool {

    ans := true

    last_a := -1
    first_b := 101
    b_found := false

    for i:=0;i<len(s);i++ {
        if s[i] == 'a' {
            last_a = i
        }
        if s[i] == 'b' && b_found == false {
            first_b = i
            b_found = true
        }

    }

    fmt.Println(last_a,first_b)

    if last_a < first_b {
        ans = true
    } else {
        ans = false
    }

    return ans

    
}
