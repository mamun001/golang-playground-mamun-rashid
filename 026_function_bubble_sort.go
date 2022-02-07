
func bubble_sort(arr []int) []int {
    L := len(arr)
    answer := make([]int, 0)
    
    // fill up answer with arr's values
    for i:=0; i < L; i++ {
        answer = append(answer,arr[i])
    }
    
    for times:=1;times<L;times++{
      for i:=0; i < L-1; i++ {
        //fmt.Println(i)
        if answer[i] > answer[i+1] {
            // out of order, lets swap the two
            temp := answer[i]
            answer[i] = answer[i+1]
            answer[i+1] = temp
        }
        //fmt.Println(answer)
      }
    }
    
    
    return answer
    
}
