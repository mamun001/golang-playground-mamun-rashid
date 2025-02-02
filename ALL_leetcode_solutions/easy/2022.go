// easy 2022 convert 1D array into
// ELO 1307
// matrix
// my alg perf: CPU > 6% and RAM > 60%
// so, almost ran out of time
// could be because I left the fmt.Println statement uncommented

func construct2DArray(original []int, m int, n int) [][]int {

    ans :=[][]int{}
    ind := -1

    // case of "impossible" as in the problem description
    if len(original) != m*n {
        return ans
    }

    for i:=0;i<m;i++{
        ans = append(ans,[]int{})
        for j:=0;j<n;j++{
          ind = n*i+j
          fmt.Println(i,j,ind,original[ind])
          ans[i] = append(ans[i],original[ind])
        }
    }
    //ans = append(ans,[]int{1,2})



    return ans
    
}

