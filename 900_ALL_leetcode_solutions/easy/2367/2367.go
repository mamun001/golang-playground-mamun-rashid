// easy 2367
// Number of Arithmetic Triplets

// 84.6 pct acc rate

// JAN 3, 2025

// I could have done it using O(n-sq) , but knew that there would be )(n)
// solution. Which turned out to be true
// Had to lookup the algorithm on youtube


// This should be classified as unicorn clever solution
// not a double pointer solution

//import "fmt"


func is_in_slice (nums []int, n int) bool {

    ans := false

    for _,v := range nums {
      if v == n {
        return true
      }
    }

    return ans
}

func arithmeticTriplets(nums []int, diff int) int {
    ans := 0

    //fmt.Println(is_in_slice(nums, 11))

    for i := len(nums) -1 ; i > -1 ; i-- {
        //fmt.Println(nums[i], is_in_slice(nums, nums[i] - diff))
        if is_in_slice(nums, nums[i] - diff) && is_in_slice(nums, nums[i] - diff*2) {
            //fmt.Println("found")
            ans++
        }
    }
    
    return ans
}
