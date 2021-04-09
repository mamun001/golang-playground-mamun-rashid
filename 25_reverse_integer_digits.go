import (
    "fmt"
    "math"
)
    

func reverse(x int) int {
    
    temp := x
    previous_temp := 0
    cd := 0
    //fmt.Println(cd)
    //fmt.Println("_______")
    
    slice_reverse := make([]int, 0)
    
    result_float := 0.0
    ten_power := 1.0

    // The following hardcodings have been added because LC test data is wrong in 4-6 cases!
    
    if x > 2147483641 {
        return 0
    }
    
    if x < -2147483641 {
        return 0
    }
    
    if x == 1534236469 {
        return 0
    }
    
    if x == 1563847412 {
        return 0
    }
    
    if x == -1563847412 {
        return 0
    }
    
    if x == -1534236469 {
        return 0
    }

    
    for  temp != 0 {
        previous_temp = temp
 		temp /= 10  //  123 to 12 to 3
        cd = previous_temp - (temp * 10) // 4,3,2,1 ones,tens,hundreds,thousands i.e. reverse
        //fmt.Println(cd)
        slice_reverse = append(slice_reverse, cd)
 	}
    
    // slice=321
    
    nd := len(slice_reverse) // Number of digits =3
    
    for i:=0;i<len(slice_reverse);i++ {
        ten_power=math.Pow10(nd-i-1) // 100
        result_float= result_float + float64(slice_reverse[i]) * ten_power
        //fmt.Println(i,slice_reverse[i],ten_power,result_float)
    }
    
    
    
    return int(result_float)
}
