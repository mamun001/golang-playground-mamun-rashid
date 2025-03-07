// easy 35, no ELO, part of Neetcode 250
// acc rate 47.8%
// part of neetcode 250



func searchInsert(nums []int, target int) int {
    l:=0
    r:=len(nums)-1
    mid := -1

    if target > nums[len(nums)-1] {
        return len(nums)
    }
    
    if target < nums[0] {
        return 0
    }

    if target == nums[len(nums)-1] {
        return len(nums)-1
    }
    
    if target == nums[0] {
        return 0
    }


    for l<r {
        mid = (l+r)/2
        //fmt.Println(mid,nums[mid])
        if target == nums[mid] {
            return mid
        } else if mid > 0 && target > nums[mid-1] && target < nums[mid]  {
            return mid
        } else if mid <len(nums)-1 && target < nums[mid+1] && target > nums[mid]  {
            return mid+1
        } else if nums[mid] > target {
            r = mid
        } else if nums[mid] < target {
            l = mid+1
        }
    }


    return -1
    
}
