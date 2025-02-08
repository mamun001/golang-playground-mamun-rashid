// medium 2150, ELO 1276
// find all lonely numbers


func findLonely(nums []int) []int {

    ans:=[]int{}

    map1:=make(map[int]int)

    for i:=0;i<len(nums);i++ {
        map1[nums[i]]++
    }

    for k,v := range map1 {
        if v == 1 && map1[k-1] == 0 && map1[k+1] == 0 {
            ans = append(ans,k)
        }
    }


    return ans
    
}

