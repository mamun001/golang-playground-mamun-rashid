// easy 1380
// Lucky Numbers in a matrix

// 79.9 acc rate

// JAN 6, 2025

import "fmt"

func luckyNumbers(matrix [][]int) []int {

    // [[1,2,3],[4,5,6]] MEANS 2 rows and 3 columns , m=2 n =3
    // 1 2 3
    // 4 5 6

    ans := []int{}


    // setup initial min and max values just outside the allowed ranges
    min_of_rows := []int{}
    max_of_columns := []int{}

    for i:=0 ; i < len(matrix) ; i++ {
        min_of_rows = append(min_of_rows,100001)
    }
    fmt.Println(min_of_rows)
    for i:=0 ; i < len(matrix[0]) ; i++ {
        max_of_columns = append(max_of_columns,0)
    }
    fmt.Println(max_of_columns)
    
    // find the mins and maxes
    for i:=0 ; i < len(matrix) ; i++ {
        for j:=0 ; j <len(matrix[0]) ; j++ {
            if matrix[i][j] > max_of_columns[j] {
                max_of_columns[j] = matrix[i][j]
            }
            if matrix[i][j] < min_of_rows[i] {
                min_of_rows[i] = matrix[i][j]
            } // if
          } // j
    } // i
    fmt.Println(min_of_rows)
    fmt.Println(max_of_columns)

    // now see which elemnt is the min_of_row AND max_of_column
    for i:=0 ; i < len(matrix) ; i++ {
        for j:=0 ; j <len(matrix[0]) ; j++ {
            if matrix[i][j] == max_of_columns[j] && matrix[i][j] == min_of_rows[i]{
                ans = append(ans,matrix[i][j])
            } // if
          } // j
    } // i

    return ans
    
}
