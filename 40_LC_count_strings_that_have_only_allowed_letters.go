// LC easy 81 pct acceptance rate

func contains(s []byte, e byte) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
} // func



func countConsistentStrings(allowed string, words []string) int {
    
    allowed_letters := make([]byte, 0)
    //consistent_or_not := true
    answer_count := 0
    this_string_ok := true

    
    for pos, char := range allowed {
       //fmt.Printf("character %c starts at byte position %d\n", char, pos)
        fmt.Println(pos)
        allowed_letters = append(allowed_letters,byte(char))
    }
    
    
    
    for pos, char := range words {
       //fmt.Printf("character %c starts at byte position %d\n", char, pos)
        fmt.Println(pos)
        this_string_ok = true
        for pos_inside, char_inside := range char {
            //fmt.Println(pos_inside,char_inside)
            fmt.Println(pos_inside)
            if contains(allowed_letters,byte(char_inside)) {
                fmt.Println("ok") 
            } else {
                //fmt.Println("BAD letter")
                this_string_ok = false         
            } // IF
            } // FOR 2
            // IF This string has survived the test for each character
        if this_string_ok == true {
            answer_count++
        }
        
        } // FOR 1
    
    return answer_count
    
    
    
} //FUNC
