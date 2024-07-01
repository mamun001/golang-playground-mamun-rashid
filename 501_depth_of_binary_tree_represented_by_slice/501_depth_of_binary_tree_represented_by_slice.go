

// Mamun Rashid code

// Tested from length 0 through 16

func depth (s1 []int) int {
    // return depth binary tree reprseneted by the slice
    // basically log of base 2 of length +1
    // horribel code, but at least my original
    L:=len(s1)
    fmt.Println("length ",L)
    answer := 0
    
    for L >= 0 {
        if L == 0 {
            answer=answer+0
            return answer
        } else if L == 1 {
            answer=answer+1
            return answer
       } else { //L is at least 2
            answer++
            L = L/2
       }
    }
    
    return answer
}
