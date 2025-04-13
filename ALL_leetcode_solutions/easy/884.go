// easy 884, ELO 1259
// uncommon words from two sentences

import "strings"

func is_in(sa []string, w string) bool {
    for i:=0;i<len(sa);i++ {
        if sa[i] == w {
            return true
        }
    }
    return false
}

func remove_all_dups(sa []string) []string{
    ans:=[]string{}
    word_map:=make(map[string]int)

    for i:=0;i<len(sa);i++{
        word_map[sa[i]]++
    }

    for k,v := range word_map{
        if v == 1 {
            ans=append(ans,k)
        }
    }
    return ans
}

func uncommonFromSentences(s1 string, s2 string) []string {

    ans:=[]string{}

    words1:=strings.Split(s1," ")
    words2:=strings.Split(s2," ")

    //fmt.Println(words1)
    //fmt.Println(words2)

    for i:=0;i<len(words1);i++{
        if is_in(words2,words1[i]) == false {
            ans=append(ans,words1[i])
        }
    }

    for i:=0;i<len(words2);i++{
        if is_in(words1,words2[i]) == false {
            ans=append(ans,words2[i])
        }
    }

    ans=remove_all_dups(ans)

    return ans
    
}
