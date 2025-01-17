// easy 1572
// matrix diagonal sum

// JAN 5, 2025

// was able to do this really fast

import "fmt"
func diagonalSum(mat [][]int) int {

ans := 0

    for i:=0 ; i < len(mat) ; i ++ {
        for j:=0 ; j <len(mat) ; j++ {
            if i == j {
                ans = ans + mat[i][j]
            } else if i+j == len(mat)-1 {
                fmt.Println(i,j)
                ans = ans + mat[i][j]
            }
        }
    }

 return ans

    
}

