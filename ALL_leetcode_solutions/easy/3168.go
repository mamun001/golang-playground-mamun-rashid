// easy 3168, ELO 1211
// Minimum number of chaors in a waiting room

func minimumChairs(s string) int {

    hwm := 0 // high water mark
    cur := 0

    for i,_ := range s {
        //fmt.Println(rune(s[i])) // 69 E 76 L
        if rune(s[i]) == 69 { // E
            //fmt.Println("E")
            cur++
            if cur > hwm {
                hwm++
            }
        } else if rune(s[i]) == 76 {
            cur--
        }
    }




    return hwm
    
}

