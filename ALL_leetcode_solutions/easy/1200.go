// easy 1200, ELO 1198
// Minimum absolute difference
// needed help from chatgpt for debugging, see blow

import "sort"


func minimumAbsDifference(arr []int) [][]int {

    ans := [][]int{}

    sort.Ints(arr)

    min_diff := 2000000 // 2 times 10 to the 6
    this_diff := 2000000
    //row := []int{0,0}

    for i:=1;i<len(arr) && min_diff != 1; i++ {
       this_diff = arr[i]-arr[i-1]
       if this_diff < min_diff {
        min_diff = this_diff
       }
    }


    fmt.Println(arr)
    for i:=1;i<len(arr); i++ {
        this_diff = arr[i]-arr[i-1]
        if this_diff == min_diff {
            row :=[]int{arr[i-1],arr[i]} // This I had to get help from chatgpt
                  // for debugging. Incorrect version has row[0]=arr[i-1] &
                  // row[1]=arr[i]. Chatgpt said that keeps reusing the old values of 
                  // row
                  // quote from chatgpt
                  // When the loop runs for the second time, it modifies row again.
                  // Since ans contains references to the same row slice, all 
                  // previous entries in ans are affected, and they all end up 
                  // with the last values assigned to row.
            ans = append(ans,row)
            //fmt.Println("ans",ans)
        }

    }

    //fmt.Println(min_diff)

    return ans
    
}
