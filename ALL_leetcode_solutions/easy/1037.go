// easy 1037, ELO 1256
// valid boomerang
// had to look up how to round a float

import "math"

func slope(x1,y1,x2,y2 int) float64{
   //if x2 == x1 {
       //return 99999.0
   //}
   return float64(y2-y1)/float64(x2-x1)
}

func intercept(x,y,m float64) float64 {
    // c = y -mx
    return y - m*x
}

func isBoomerang(points [][]int) bool {


    // first and second point same
    if points[0][0] == points[1][0] && points[0][1] == points[1][1] {
        return false
    }

    // 2nd and 3rd point same
    if points[1][0] == points[2][0] && points[1][1] == points[2][1] {
        return false
    }

    // 3rd and first point same
    if points[2][0] == points[0][0] && points[2][1] == points[0][1] {
        return false
    }

    // all 3 Xs the same
    if points[0][0] == points[1][0] && points[0][0] == points[2][0] {
        return false
    }

    // all 3 Yss the same
    if points[0][1] == points[1][1] && points[0][1] == points[2][1] {
        return false
    }



    // y = mx + c
    // c = y - mx

    // test the func
    //fmt.Println(slope(1,7,4,9))

    m1 := slope(points[0][0],points[0][1],points[1][0],points[1][1])
    m2 := slope(points[1][0],points[1][1],points[2][0],points[2][1])
    //fmt.Println(m1,m2) 
    // without rounding we get false "not equal" because precision problems when comparing floats
    i1 := math.Round(intercept(float64(points[0][0]),float64(points[0][1]),m1))
    i2 := math.Round(intercept(float64(points[2][0]),float64(points[2][1]),m2))
    //fmt.Println(i1,i2)

    return m1 != m2 || i1 != i2
    
}
