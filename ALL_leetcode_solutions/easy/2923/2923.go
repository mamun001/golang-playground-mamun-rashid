// DUMBEST explanation of a problem!!

// grid is W table in football tournament!
// "1" means "W" against corresponding "column"

// ALL WE DO is see which row has most 1s just like in round robin tournament
// "points"

// easy 2923
// find champion I

// JAN 8, 2025



func findChampion(grid [][]int) int {

    ones_in_rows := []int{} // we count 1s in each row
    n := len(grid)
    max_ones := 0 // max number of ones
    winner_row := 999999 // which row has most ones?

    // give initial values
    for i := 0 ; i < n ; i++ {
        ones_in_rows = append (ones_in_rows,0)
    }

    // count 1s in each row
    for r := 0 ; r < n ; r++ {
        for c:= 0 ; c < n ; c++ {
            if grid[r][c] == 1 {
                ones_in_rows[r]++
            }
        }
    }

    // which row has the max number of ones?
    for r := 0 ; r < n ; r++ {
        if ones_in_rows[r] > max_ones {
            max_ones = ones_in_rows[r]
            winner_row = r

        }
    }

   
    return winner_row
   
}
