// easy 2574
// left and right sum diff
// JAN 13, 2025
// my alg: cpu beats 100%, RAM beats only 9 pct


import "fmt"

func abs (n int) int {
    ans := 0
    if n < 0 {
        ans = n * -1
    } else {
        ans = n + 0
    }

    return ans
}

func leftRightDifference(nums []int) []int {
    ans := []int{}
    L :=  len (nums)
    left_arr := []int {}
    right_arr := []int {}

    // initialize all 3 arrays
    for i:=0 ; i < L ; i ++ {
        left_arr = append(left_arr,0)
        right_arr = append(right_arr,0)
        ans = append(ans,0)
    }

    // figure out left sums
    for i:=0 ; i < L ; i++ {
        for j:=0 ; j < i ; j++ {
           //fmt.Println(nums[j])
           left_arr[i] = left_arr[i] + nums[j]
        }
        //fmt.Println("end of i")
    }
    //fmt.Println(left_arr)

    // figure our right sums
    for i:=0 ; i < L ; i++ {
        for j:=i+1 ; j < L ; j++ {
           //fmt.Println(nums[j])
           right_arr[i] = right_arr[i] + nums[j]
        }
        //fmt.Println("end of i")
    }
    //fmt.Println(right_arr)

    // figure out the ans array
    for i:=0 ; i < L ; i++ {
        ans[i] = abs ( left_arr[i] - right_arr[i])
    }
    return ans
    
}

