// 3289 The two sneaky numbers of digitville

// acc rate 89.1%

// Dec 23 , 2024


import "fmt"


// we need a func that returns true if a slice has a given number in it
func in_slice (s []int, n int) bool {
    
    ans := false 
    for _,v := range s {
       if v == n {
        ans = true
       }
    }

    return ans
}

func getSneakyNumbers(nums []int) []int {
    //test := in_slice (nums, 9)

    // we insert into holding_slice if the int we are looking at doesn't already have that int in it
    holding_slice := []int{}
    ans := []int{}

    for _,v := range nums {
        if in_slice (holding_slice, v) {
            ans = append(ans,v) // found a duplciate!
        } else {
            holding_slice = append(holding_slice,v) // this int is a new one
        }
    }

    //fmt.Println(test)

    return ans
}
