// easy 830,, ELO 1252
// positions of large groups

func largeGroupPositions(s string) [][]int {

    ans:=[][]int{}

    cur_count := 1

    for i:=1;i<len(s);i++ {
        if s[i] == s[i-1] { // streak continues
            cur_count++
        } else {
            // streak just broke
            if cur_count > 2 {
                //fmt.Println(i-cur_count,i-1)
                temp:=[]int{i-cur_count,i-1}
                ans = append(ans,temp)
            }
            cur_count = 1
        }
    }

    // if the last chars are still on a streak
    if cur_count > 2 {
                //fmt.Println(i-cur_count,i-1)
                temp:=[]int{len(s)-cur_count,len(s)-1}
                ans = append(ans,temp)
            }



    return ans
    
}
