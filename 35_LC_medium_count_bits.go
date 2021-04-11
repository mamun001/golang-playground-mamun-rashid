func countBits(num int) []int {
    
    answer := make([]int, 0)
    
    for i:=0;i<=num;i++ {
      s := fmt.Sprintf("%b", i)
      //fmt.Println(s)
        
        count_of_1s := 0
        for pos, char := range s {
           //fmt.Printf("character %c starts at byte position %d\n", char, pos)
            fmt.Println(pos)
            if char == '1' {
                count_of_1s++
            }
        } // FOR
        answer = append (answer,count_of_1s)
    }
      
      
    
    return answer
    
}
