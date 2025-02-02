// 2319 easy check if matrix is x-matrix
// ELO 1201


func checkXMatrix(grid [][]int) bool {

    ans := true

    n := len(grid)

    // first diagonal
    for i:=0; i<n ; i++ {
        if grid[i][i] == 0 {
            //fmt.Println(i,i)
            return false
        }
    }


    // 2nd diagonal
    for i:=0; i<n ; i++ {
        for j:=0; j<n ; j++ {
            if i+j == n-1 && grid[i][j] == 0  {
                 return false
            }
        }
    }
    

    // non-diagonals
    for i:=0; i<n ; i++ {
        for j:=0; j<n ; j++ {
            if i == j || i+j == n-1 {
                 i = i // do nothing, dont check
            } else {
                if grid[i][j] != 0 {
                    //fmt.Println(i,j)
                    return false
                }
            }

        }
    }




    return ans
    
}

