// easy 1232 , ELO 1247
// Check if it is a straight line

func checkStraightLine(coordinates [][]int) bool {

    xdiff := coordinates[1][0] - coordinates[0][0]
    ydiff := coordinates[1][1] - coordinates[0][1]
    slope := float64(ydiff)/float64(xdiff)

    fmt.Println(xdiff)
    fmt.Println(ydiff)
    fmt.Println(slope)


    // keep checking for slope equality
    
    for i:=2;i<len(coordinates);i++ {
        cur_xdiff := coordinates[i][0] - coordinates[i-1][0]
        cur_ydiff := coordinates[i][1] - coordinates[i-1][1]
        cur_slope := float64(cur_ydiff)/float64(cur_xdiff)
        fmt.Println(cur_xdiff,cur_ydiff,cur_slope)
        if xdiff == 0 && cur_xdiff == 0 { // when slope infinity, it is an edge case
            continue
        }
        if cur_slope != slope  {
            return false
        }
    }
    



    return true
    
}
