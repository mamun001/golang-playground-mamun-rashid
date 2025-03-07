// easy 3042 Count Prefix and Suffix Pairs I
// ELO 1214


func isPrefixAndSuffix(s1, s2 string) bool {

    n1 := len(s1)
    n2 := len(s2)

    if n1 > n2 {
        return false
    }

    //fmt.Println(s2[:n1],s2[n2-n1:])

    if s1 == s2[:n1] && s1 == s2[n2-n1:] {
        return true
    }

    return false

}

func countPrefixSuffixPairs(words []string) int {

    ans:=0

    //s1 := "a"
    //s2 := "abcabad"
    //fmt.Println(isPrefixAndSuffix(s1,s2))

    for i:=0;i<len(words);i++ {
        for j:=0;j<len(words);j++ {
            if i < j && isPrefixAndSuffix(words[i],words[j]) {
               //fmt.Println(words[i],words[j])
               //fmt.Println(isPrefixAndSuffix(words[i],words[j]))
               ans++
        }
        }
    }


    return ans
    
}
