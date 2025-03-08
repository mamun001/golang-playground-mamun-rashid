// easy, 1422, ELO 1238
// Maximum Score After Splitting

func score(s1,s2 string) int {
    ans:= 0

    for i:=0;i<len(s1);i++ {
        if s1[i] == '0' {
        ans++
        }
    }

    for i:=0;i<len(s2);i++ {
        if s2[i] == '1' {
        ans++
        }
    }

    return ans

}

func maxScore(s string) int {


    //fmt.Println(score("",s))
    l:=len(s)
    max:=0

    for i:=1;i<l;i++ {
        left := s[:i]
        right := s[i:]
        //fmt.Println(score(left,right))
        if score(left,right) > max {
            max = score(left,right)
        }

    }



    return max
    
}
