// easy 883
// Projection area of 3D shapes
// description should have included more clarity and more detail, but whatever

// my algo:
// top = count on non-zero items
// front = sum of row maxes
// side = sum of col maxes


// my code's perf: cpu beats 100%, RAM beats: 43%


import "fmt"

func projectionArea(grid [][]int) int {

    ans := 0
    n := len(grid)
    top := 0
    row_max_sum := 0 // "front"
    col_max_sum := 0 // "side"
    row_max := []int{}
    col_max := []int{}
    // initialize with zeros
    for r:= 0 ; r < n ; r++ {
        row_max = append (row_max,0)
    }
    for c:= 0 ; c < n ; c++ {
        col_max = append (col_max,0)
    }

    // find row maxes, columns and count non-zero items
    for r:= 0 ; r < n ; r++ {
        for c:= 0 ; c < n ; c++ {
            if grid[r][c] > row_max[r] { // "front"
                row_max[r] = grid[r][c]
            }
            if grid[r][c] > col_max[c] { // "side"
                col_max[c] = grid[r][c] 
            }
            if grid[r][c] > 0 { // if height is there is no shadow from the top
               top++
            }
        }
    }

    // get "front"
    for r:= 0 ; r < n ; r++ {
        row_max_sum = row_max_sum + row_max[r]
    }

    // get "side"
    for c:= 0 ; c < n ; c++ {
        col_max_sum = col_max_sum + col_max[c]
    }

    fmt.Println(top,row_max_sum,col_max_sum)
    ans = top + row_max_sum + col_max_sum
    return ans
    
}

