// easy 2057, ELO 1167
// Smallest index with equal value

// JAN 20, 2025

func smallestEqual(nums []int) int {

    ans := -1

    for i:=0 ; i < len(nums); i++ {
        if i % 10 == nums[i] {
            return i
        }
    }

    return ans
    
}
