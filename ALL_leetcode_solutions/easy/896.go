// easy 896, ELO 1259
// monotonic array

func isMonotonic(nums []int) bool {

    hint:=""

    for i:=0;i<len(nums)-1;i++ {
        if nums[i+1]>nums[i] {
            hint="increasing"
            break
        }
        if nums[i+1]<nums[i] {
            hint="decreasing"
            break
        }
    }

    if hint=="" { // all items are equal
        return true
    }

    if hint == "increasing" {
        for i:=0;i<len(nums)-1;i++ {
          if nums[i+1]<nums[i] {
              return false
          }
        }
    }

    if hint == "decreasing" {
        for i:=0;i<len(nums)-1;i++ {
          if nums[i+1]>nums[i] {
              return false
          }
        }
    }

    return true
    
}
