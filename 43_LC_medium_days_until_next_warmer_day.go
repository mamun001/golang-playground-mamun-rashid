
// https://leetcode.com/problems/daily-temperatures/

func dailyTemperatures(T []int) []int {
    answer := make([]int, 0)
    var wait int = 99999
    
    for i:=0; i<len(T); i++ {
        if i== len(T) -1 {
            answer = append(answer,0) // last day will never see a warmer day
        }
        for j:=i+1; j<len(T); j++ {
            if T[j] > T[i] {
                wait = j-i
                //fmt.Println(i,T[i],j,T[j],wait)
                answer = append(answer,wait)
                break
            } else if j == len(T)-1 {
                answer = append(answer,0)
            } // else if
        } // for in 
    } // for out
    
    return answer
} // FUNC
