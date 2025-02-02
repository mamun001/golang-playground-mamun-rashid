// easy 3402, minimum operations to make columns
// ELO unknown
// acc rate 73%

func minimumOperations(grid [][]int) int {

    m := len(grid)
    n := len(grid[0])

    needed :=0
    highc := 0
    diff := 0

    for j:=0;j<n;j++{
        //fmt.Println("new colum........")
        highc = grid[0][j]
        //fmt.Println(highc)
        for i:=0;i<m;i++{
            //fmt.Println(grid[i][j])
            if i == 0 {
                i = i // do nothing
            } else if grid[i][j] > highc {
                //fmt.Println("already higher")
                highc = grid[i][j]
            } else {
                //fmt.Println("lower or equal)")
                diff = highc - grid[i][j]
                needed = needed + diff + 1
                highc = highc + 1
                //fmt.Println("need",needed)
            }
        }
        //fmt.Println(needed,highc)
    }
    return needed
    
}

