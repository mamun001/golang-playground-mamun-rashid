// easy 3105, ELO 1217
// Longest Strictly Increasing or Strictly

func longestMonotonicSubarray(nums []int) int {

    cur_asc := 0
    cur_des := 0
    max_asc := 0
    max_des := 0

    for i:=1; i<len(nums);i++ {
       if nums[i] == nums[i-1] {
           cur_asc = 0
           cur_des = 0
       } else if nums[i] > nums[i-1] {
           cur_asc++
           cur_des = 0
           if max_asc < cur_asc {
               max_asc = cur_asc
           }
       } else if nums[i] < nums[i-1] {
           cur_des++
           cur_asc = 0
           if max_des < cur_des {
               max_des = cur_des
           }
       }
    }

    if max_asc > max_des {
        return max_asc+1
    } else {
        return max_des+1
    }
    
}
