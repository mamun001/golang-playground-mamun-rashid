// MEDIUM, 560, subarray of sum equals k
// hardest one form neetcode's 250 for array and hasmap set
// Got lots of help from youtube and chatgpt for this one
// THIS ONE NEEDS TO PRACTICED 10 Times

// My alg worked just fine for all positive numbers (two pointers)
// But will not work when negative numbers are involved

func subarraySum(nums []int, k int) int {

    ans := 0
    prefix_sum := 0 // This one will build and build from left to right
    prefix_map := make(map[int]int) // This is where we will keep track of frequencies of each prefix sum
    prefix_map[0] = 1 // base case , why ????

    for i:=0 ; i<len(nums); i++ {
        prefix_sum = prefix_sum + nums[i] // this one will grow and grow left to right

        // THIS IS HEART OF THiS FANCY ALGORITHM
        // if there is entry in the map for prefix_sum -k 
        // then we increment ans
        val,exists := prefix_map[prefix_sum-k] 
        if exists == true {
            ans = ans + val // THIS I DONT UNDERSTAND
        }
        prefix_map[prefix_sum]++ // we have found a new occurance of this "prefix_sum"
    } 

    return ans


    return ans
    
}
