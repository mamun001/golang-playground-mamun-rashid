// easy 1160 , ELO 1205

func countCharacters(words []string, chars string) int {

    ans := 0
    map1 := make(map[byte]int)

    for i:=0;i<len(chars);i++ {
        map1[chars[i]]++
    }
    //fmt.Println(map1)
    //fmt.Println(" ")

    //map2 := make(map[byte]int)

    for i:=0;i<len(words);i++ {
       map2 := make(map[byte]int)
       for j:=0;j<len(words[i]);j++ {  
           //fmt.Println(j)      
           map2[words[i][j]]=0
        }
       //fmt.Println(map2)
       //fmt.Println(words[i])
       for j:=0;j<len(words[i]);j++ {
           //fmt.Println(j)
           //fmt.Println(words[i][j])
           map2[words[i][j]]++
        }
        // now check that for each key, the value in map1 => more what we need
        poss := true
        
        //fmt.Println(words[i])
        //fmt.Println(map2)
        for k,_ := range map2 {
            if map1[k] < map2[k] {
                poss = false
            }
        }
        //fmt.Println(words[i],poss)
        if poss == true {
            ans = ans + len(words[i])
        }
    }


    return ans
    
}
