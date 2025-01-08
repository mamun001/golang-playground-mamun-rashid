// easy 2500
// 78.6 acc rate

// Delete Greatest Value in Each Row

// JAN 7, 2025

// HAD TO LOOKUP ALGORITHM FROM YOUTUBE
// ALSO HOW TO SORT SLICE ROWS IN PLACE


import "fmt"
import "sort"


func deleteGreatestValue(grid [][]int) int {

    ans := 0
    m := len(grid) // rows
    n := len(grid[0]) // columns

    max_of_cols := []int{} // starting values
    for c := 0 ; c < n ; c++ {
        max_of_cols = append(max_of_cols,1)
    }
    fmt.Println(max_of_cols)


     //proving that sort.Ints sorts in place
    //fmt.Println(grid[0])
    //sort.Ints(grid[0])
    //fmt.Println(grid[0])


    // sort each row
    for r := 0; r < m ; r++ {
        sort.Ints(grid[r])
    }

    // find max of each col
    for r := 0; r < m ; r++ {
        for c := 0 ; c < n ; c++ {
           //fmt.Println(grid[r][c])
           if grid[r][c] > max_of_cols[c] {
            max_of_cols[c] = grid[r][c]
           } // if
        } // inner for
        
        //ans = ans + grid[r][n-1]
    } // outer for

    //fmt.Println(max_of_cols)

    // add max of each col
    for c := 0 ; c < n ; c++ {
        ans = ans + max_of_cols[c]
    }

    return ans
    
}
