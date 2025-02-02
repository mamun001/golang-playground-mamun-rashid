// easy 867 Transpose Matrix, ELO 1258
// 

func transpose(matrix [][]int) [][]int {

    ans := [][]int{}

    m := len(matrix)
    n := len(matrix[0])

    // rows becomes columns, columns becomes rows

    for i:=0;i<n;i++ {
        ans = append(ans,[]int{})
        for j:=0;j<m;j++{
            ans[i] = append(ans[i],matrix[j][i])
        }
    }

    return ans
    
}

