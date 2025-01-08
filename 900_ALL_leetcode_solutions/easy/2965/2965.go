// easy 2965
// 76.2 pct acc rate
// find missing and repeated values

// JAN 7, 2025
// CAME UP WITH THIS MATH BASED ALGORITHM ON MY OWN
//    finding the dup is easy
// missing = diff_bw_expected_sum_and_actual_sum - duplicate_number



import "fmt"

func is_in_slice(nums []int , m int) bool {
    ans := false
    for i:=0; i < len(nums); i++ {
        if nums[i] == m {
            return true
        }
    }
    return ans
}

func findMissingAndRepeatedValues(grid [][]int) []int {
   ans := []int {}
   n := len(grid)
   //fmt.Println(n)
   actual_sum := 0
   expected_sum := (n*n)*(n*n+1)/2
   found_nums := []int{}
   found_dup := 0 // fake start value
   missing := 0 // fake start value

   // test the func
   //s1 := []int {1,2,3}
   //fmt.Println(is_in_slice(s1,2))
   //fmt.Println(is_in_slice(s1,4))

   // do ONE run and find the dup and actual sum
   for i:=0 ; i < n ; i++ {
       for j:=0 ; j < n ; j++ {
        actual_sum = actual_sum + grid[i][j]
        if is_in_slice(found_nums,grid[i][j]) { // we this number already, so it is the dup
           found_dup = grid[i][j]
        } else {
            found_nums = append(found_nums,grid[i][j])
        }
       }
   }

   // formula I came up with
   missing = found_dup - actual_sum + expected_sum 


   //fmt.Println(expected_sum, actual_sum)
   //fmt.Println(found_dup)
   //fmt.Println(missing)

   ans = append(ans,found_dup)
   ans = append(ans,missing)
    
   return ans
}
