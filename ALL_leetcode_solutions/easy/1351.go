


// leetcode easy 1351
// nested_loop

func countNegatives(grid [][]int) int {
    count := 0
    
    for i:=0; i<len(grid); i++ {
      for j:=0; j<len(grid[i]); j++ {
          //fmt.Println(grid[i][j])
          if grid[i][j] < 0 {
              count++
          }
      }       
    }
    
    return count
} // FUNC
