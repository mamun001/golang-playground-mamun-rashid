// easy 3120 , ELO 1206
// count the Number of Special Characters

func numberOfSpecialChars(word string) int {

   ans:= 0

    // 65-90
    // 97-122
    //fmt.Println(byte('A'))
    //fmt.Println(byte('Z'))
    //fmt.Println(byte('a'))
    //fmt.Println(byte('z'))

    map1:=make(map[byte]int)

    for i:=0;i<len(word);i++ {
        map1[byte(word[i])]++
    }
    //fmt.Println(map1)

    for i:=65;i<91;i++ {
        if map1[byte(i)] > 0 && map1[byte(i+32)] > 0 {
            ans++
        }
    }


    return ans
    
}
