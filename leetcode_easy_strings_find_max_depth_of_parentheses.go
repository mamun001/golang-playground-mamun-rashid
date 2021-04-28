

// leetcode easy
// strings

func maxDepth(s string) int {
    
    
    diff_open_close :=0
    max := 0
    
    for pos, char := range s {
        if char == '(' {
            fmt.Println(pos) // just so pos gets used and golang compiler just not complain
            diff_open_close++  
            //fmt.Println(diff_open_close)
            if diff_open_close > max {
               max = diff_open_close
            }
        } else if char == ')' {
            //fmt.Printf("character %c starts at byte position %d\n", char, pos)
            diff_open_close--
            //fmt.Println(diff_open_close)
        }
        
    }
    
    return max
}

