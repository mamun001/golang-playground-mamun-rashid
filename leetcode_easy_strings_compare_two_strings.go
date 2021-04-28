
// leetcode easy 
// strings


func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    
    combo_1 := make([]byte, 0)
    combo_2 := make([]byte, 0)
    answer := true
    
    
    
    for i:=0;i<len(word1);i++{
        //fmt.Println(i)
        for pos, char := range word1[i] {
            //fmt.Printf("character %c starts at byte position %d\n", char, pos)
            fmt.Println(pos)
            combo_1 = append(combo_1,byte(char))
        }
    }
    
    
    for i:=0;i<len(word2);i++{
        //fmt.Println(i)
        for pos, char := range word2[i] {
            //fmt.Printf("character %c starts at byte position %d\n", char, pos)
            fmt.Println(pos)
            combo_2 = append(combo_2,byte(char))
        }
    }
    
    //fmt.Println("----------")
    
    if len(combo_1) != len(combo_2) {
        answer = false
        return answer
    } else {
      for i:=0;i<len(combo_1); i++ {
          if combo_1[i] != combo_2[i] {
              answer = false
              return false
          } // IF
      } // FOR
    } // IF
        
    return answer
}
