// 3194 Minimum average of smallest and largest elements

// acc rate 84.9%

// DEC 24, 2024
// GOT HELP from youtube ****

import "fmt"

func min (floats []float64) float64 { // returns the minimum in an array
    ans := floats[0]

    for _,v := range floats {
        if v < ans {
            ans = v
        }
    }

    return ans
 
}

func minimumAverage(nums []int) float64 {

    
    averages := []float64{} // as the question asked for
    current_average := 0.0 // for each pair in the loop

    sort.Ints(nums) // first we sort the given array
    // two pointers algorithm
    a:=0
    z:=len(nums)-1 // last one

    // how many times we run the loop
    n_by_2 := len(nums)/2 

    for i := 0 ; i < n_by_2  ; i++ {
        //fmt.Println(i,a,z,nums[a],nums[z])
        current_average = float64(nums[a]) / 2.0 + float64(nums[z]) / 2.0
        //fmt.Println(current_average)
        averages = append (averages, current_average) // add to it the "averages" array
        a++ // move the pointer
        z-- // move the pointer
    }

    return min(averages)
}

