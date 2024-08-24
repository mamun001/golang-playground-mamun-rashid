
// Coder: Mamun Rashid

func reverse(s1 []int) []int {
    answer:= make([]int,0)
    
    
    L := len(s1)
    for i:=L-1; i >=0 ; i-- {
        answer=append(answer,s1[i])
    }
    
    return answer
}

