// easy 2395 , ELO 1250
// Find subarrays with equal sum

func findSubarrays(nums []int) bool {

    map_sums := make(map[int]int)

    for i:=0;i<len(nums)-1;i++ {
        sum := nums[i]+nums[i+1]
        _,exists := map_sums[sum]
        if exists {
            return true
        } else {
            map_sums[sum]++
        }
    }


    return false
    
}
