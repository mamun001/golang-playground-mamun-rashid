// easy 3033, ELO 1180
// modify the matrix
// Jan 23, 2025

func modifiedMatrix(matrix [][]int) [][]int {

    answer := matrix
    row := len(matrix)
    col := len(matrix[0])
    col_max := make([]int,col)


    //answer[0][0] = 999
    //fmt.Println(matrix[0][0])
    //fmt.Println(answer[0][0])

    // find the column maxes
    for i:=0;i<row;i++ {
        for j:=0;j<col;j++{
            if matrix[i][j] > col_max[j] {
                col_max[j] = matrix[i][j]
            }
        }
    }
    //fmt.Println(col_max)

    // replace -1 with column maxes
    for i:=0;i<row;i++ {
        for j:=0;j<col;j++{
            if matrix[i][j] == -1 {
                answer[i][j] = col_max[j]
            }
        }
    }


    return answer
    
}
