// 1941 easy ELO 1242 
// check if all characters have equal
// a classic array/hashmap problem

func areOccurrencesEqual(s string) bool {

    ans := true
    map1 :=make(map[rune]int)


    for i,_ := range s {

        map1[rune(s[i])]++
    }


    //fmt.Println(map1)

    first_count := map1[rune(s[0])]
    for _,v := range map1 {
        if v != first_count {
            return false
        }
    }


    return ans
    
}

