// MEDIUM 215, kth largest element in an array
// ELO unknown
// no need for heap! simply using sort library was good enough
// my alg: CPU > 93% RAM > 75%
// so, ya go library to the rescue
// this probably means go's sort is faster than Python's heap

import "sort"

func findKthLargest(nums []int, k int) int {

    sort.Ints(nums)
    n := len(nums)
    ans := nums[n-k]

    return ans
    
}

