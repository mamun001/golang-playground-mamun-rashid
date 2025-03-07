// easy

// 87.8 pct

// Largest Local Values in a Matrix

// FOR THIS I had to learn how to implement
// 2 dimensional arrays using slice of slice


import "fmt"

func largestLocal(grid [][]int) [][]int {

    var ans [][]int
    //var temp [][]int

    // how to add 1 item at time, new row or not
    //ans = append(ans, []int{})
    //ans[0] = append (ans[0],1)
    //ans[0] = append (ans[0],2)
    //ans = append(ans, []int{})
    //ans[1] = append (ans[1],3)
    //ans[1] = append (ans[0],4)
 

    HW := len(grid) -2  // height = length = of matrix
    
    // print the whole matrix
    //for i:=0; i < len(grid) ; i++ {
        //or j:=0; j<len(grid[i]); j++ { 
            //fmt.Println(grid[i][j])
        //}
    //}
    
    for i:=0; i < HW ; i++ {
        for j:=0; j < HW; j++ {
            if j == 0 { // first in row
               ans = append(ans, []int{}) // add a new empty row
            }
            // this is where we need to find the local max of always 9 items
            // i+1,j+1 = index in GRID of the exact item for whom we are trying to find max for, max cannot be 
            // less than 1 or more than 100 per the contstraists give
            max := 0 // to start with
            for k:=i; k < i+3 ; k++ { // exactly 3 times
               for l:=j; l < j+3 ; l++ {// exactly 3 times
                   //fmt.Println(grid[k][l])
                   if grid[k][l] > max {
                    max = grid[k][l]
                   }
               } // inner max loop
            } // outer max loop
            //fmt.Println(max)
            //ans[i] = append (ans[i], grid[i+1][j+1])  // insert an item on the ith row
            ans[i] = append (ans[i], max)
            //fmt.Println(grid[i+1][j+1])
        } // inner loop
    } // outer loop

    return ans
} // func

