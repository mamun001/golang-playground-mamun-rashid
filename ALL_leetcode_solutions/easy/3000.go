// easy 3000, ELO 1250
// maximum area of longest diagnoal

import "math"

func areaOfMaxDiagonal(dimensions [][]int) int {


    //fmt.Println(math.Sqrt(10))

    max_diag:=0.0
    max_area:=0

    // area := dimensions[i][0] * dimensions[i][1]

    // first loop to get max_diag
    for i:=0;i<len(dimensions);i++ {
        diag := math.Sqrt(float64(dimensions[i][0]*dimensions[i][0] + dimensions[i][1]*dimensions[i][1]))
        //fmt.Println(i,diag)
        if diag > max_diag {
            max_diag = diag
        }
    }
    //fmt.Println(max_diag)

    // 2nd loop to find max area among the max_diag's
    // because it is possible have a bigger are with smaller diagoanl; such test cases exist
    for i:=0;i<len(dimensions);i++ {
        diag := math.Sqrt(float64(dimensions[i][0]*dimensions[i][0] + dimensions[i][1]*dimensions[i][1]))
        if diag == max_diag {
            area := dimensions[i][0] * dimensions[i][1]
            if area > max_area {
                max_area = area
            }
        }
    }

    return max_area
    
}
