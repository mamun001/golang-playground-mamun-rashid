// easy 3065 Minimum Operations to

// ELO 1149

// easy algorithm

// JAN 16, 2025

func minOperations(nums []int, k int) int {
    ans :=0
    for i:=0 ; i < len(nums); i++ {
      if nums[i] < k {
        ans++
      }
    }

    return ans
}

