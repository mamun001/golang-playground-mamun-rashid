// easy 3210 find the encrypted string
// ELO 1179

// JAN 23, 2025

func getEncryptedString(s string, k int) string {

    ans := ""
    j := 0

    for i:=0; i<len(s); i++ {
        j = (i+k) % len(s) // index of the letter that will go into the answer string
        //fmt.Println(i,j,string(s[j]))
        ans = ans + string(s[j])

    }

    return ans
    
}
