// easy 2341 ELO 1184
// maximum number of pairs in array

// Jan 23, 2025

// fancy description, but it comes down to:
// how many pairs and how singles



func numberOfPairs(nums []int) []int {

    ans := []int{0,0}
    pairs := 0
    hashmap := make(map[int]int)

    // count freq of each number
    for i:=0 ; i< len(nums); i++ {
        hashmap[nums[i]]++
    }

    // count pairs
    for _,v := range (hashmap) {
        pairs = pairs + v / 2
    }
    
    ans[0] = pairs
    ans[1] = len(nums) - (pairs * 2)

    return ans
}

