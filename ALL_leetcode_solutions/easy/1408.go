/*

easy 1408, ELO 1203
string matching in an array

*/


import "strings"

func is_in(ar []string, s string) bool {
    ans := false
    for i:=0;i<len(ar);i++ {
        if s == ar[i] {
            return true
        }
    }
    return ans
}

func stringMatching(words []string) []string {

    ans:=[]string{}

    for i:=0;i<len(words);i++{
        for j:=0;j<len(words);j++{
            if i != j && strings.Index(words[i],words[j]) != -1 && is_in(ans,words[j]) == false{
                ans = append(ans,words[j])
            }
        }
    }


    return ans
    
}
