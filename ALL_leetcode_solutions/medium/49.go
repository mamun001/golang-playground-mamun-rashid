// MEDIUM 49 , no ELO
// #6 of 22 in neetcode's array and hashmap list
// NEEDED LOTS OF HELP from Youtube and chatgpt
// my alg did not work, youtube's alg did
// had to learn/copy how to sort a slice of runes from chatgpt
// had to learn how to define a map whose value is []strings from chatgpt

// THIS ONE NEEDS TO PRACTICED 5 Times

import "sort"

func sortword (s string) string {
    runes := []rune(s)

    sort.Slice(runes, func(i, j int) bool {return runes[i] < runes[j]})

    return string(runes)
}

func is_in(ar []string, s string) bool {
    ans := false
    for i:=0; i<len(ar);i++ {
        if s == ar[i] {
            return true
        }
    }
    return ans
}

func groupAnagrams(strs []string) [][]string {

    ans :=[][]string{}

    //s := "baakfhakshfajsfhajfsgahqyrqiy"
    //fmt.Println(sortword(s))

    map_anagram := make(map[string][]string)
    //fmt.Println(map_anag)
    //map_anagram[sortword("ahfjah")] = []string{"ab"}
    //fmt.Println(map_anagram)
    //map_anagram[sortword("ahfjah")] = append(map_anag[sortword("ahfjah")],"ba")
    //fmt.Println(map_anagram)

    for i:=0;i<len(strs);i++ {
        map_anagram[sortword(strs[i])] = append(map_anagram[sortword(strs[i])],strs[i])
    }
    //fmt.Println(map_anag)

    for _,v := range map_anagram {
        ans = append(ans,v)
    }                 

    return ans
    
}
