// easy 1961, ELO 1234
// check if string is a prefix of array

func isPrefixString(s string, words []string) bool {

    l := len(s)
    lw := 0
    temp_s := ""

    for i:=0;lw<=l && i<len(words);i++ {
        lw = lw + len(words[i])
        temp_s = temp_s + words[i]
        if lw == l && temp_s == s {
            return true
        }
    }



    return false
    
}
