// easy 2357, ELO 1225
// Make Array Zero by Subtracting

import "sort"

func get_first_non_zero(ar []int) int {
    // assumes ar is sorted
    for i:=0;i<len(ar);i++ {
        if ar[i] > 0 {
            return ar[i]
        }
    }
    return -1
}

func minimumOperations(nums []int) int {

    sort.Ints(nums)
    max := nums[len(nums)-1]
    //fmt.Println(max)

    if max == 0 {
        return 0
    }

    //fmt.Println(get_first_non_zero(nums[3:]))

    
    count := 0
    for true { // infinite loop
      count++
      nz := get_first_non_zero(nums)
      for i:=0;i<len(nums);i++ {   
          if nums[i] > 0 {
             nums[i] = nums[i] - nz
          }     
      }
      if nums[len(nums)-1] == 0 {
          return count
      }
    }
    return -1
    
}
