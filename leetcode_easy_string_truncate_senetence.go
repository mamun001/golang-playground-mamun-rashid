// leetcode easy #1816 Truncate Sentence
// String

func truncateSentence(s string, k int) string {
    
    answer := ""
    word_no := 1
    
    // we keep count which word we are at (1st, 2nd etc)
    // if we are at kth word or less, we keep adding adding characters to the answer string
    // as soon as we start k+1 th word, we quit traversing and return
    
    for _, char := range s {
       //fmt.Printf("character %c starts at byte position %d\n", char, pos)
        if string(char) == " " {
            //fmt.Println("space")
            word_no++
            fmt.Println(word_no)
            if word_no > k {
                fmt.Println("time to truncate")
                break
            }
        } // if string(char) ==
        answer = answer + string(char)
    }
    
    return answer
    
}
