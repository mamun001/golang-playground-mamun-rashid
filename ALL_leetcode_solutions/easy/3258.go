// easy 3258, ELO 1258
// Count Substrings That Satisfy K-constraint I

func satisfies(s string, k int) bool {
    ones:=0
    zeros:=0

    for i:=0;i<len(s);i++ {
        if s[i] == '0' {
            zeros++
        }
        if s[i] == '1' {
            ones++
        }
    }

    if zeros <= k || ones <= k {
        return true
    }
    return false
}


func countKConstraintSubstrings(s string, k int) int {

    //testing
    //fmt.Println(satisfies(s,3))

    
    L:=len(s) // length of s
    //ssl:=3 // substring length

    ans:=0
    for ssl:=1;ssl<=L;ssl++{
      for j:=0;j<=L-ssl;j++{
           substring:=s[j:j+ssl] // creating substring
           //fmt.Println(substring)
           //fmt.Println(satisfies(substring,k))
           if satisfies(substring,k) {
            ans++
           }
      }
    }

    return ans
    
}
