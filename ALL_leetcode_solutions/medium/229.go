// MEDIUM 229, majority element II
// no ELO 
// Part of Neetcode 250
// amy alg beats CPU > 100
//         beats RAM > 94%

func majorityElement(nums []int) []int {

    target := len(nums)/3 + 1
    ans := []int{}

    map1 := make(map[int]int)

    for i:=0;i<len(nums);i++ {
        map1[nums[i]]++
    }

    for k,v := range map1 {
        if v >= target {
            ans = append(ans,k)
        }
    }


    return ans
    
}
