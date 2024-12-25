// 1913 maximum product diff bw two pairs

// 82.9 percet

// DEC 25, 2024

import "fmt"

func maxProductDifference(nums []int) int {

    sort.Ints(nums) // sort first
    ////fmt.Println(nums)
    // multiply two biggest and multiply the two smallest
    ans := nums[len(nums)-1] * nums[len(nums)-2] - nums[0] * nums[1]

    return ans
    
}

