// easy 1748, ELO 1228
// Sum of Unique Elements

func sumOfUnique(nums []int) int {

    map1 := make(map[int]int)

    for i:=0;i<len(nums);i++ {
        map1[nums[i]]++
    }

    sum:=0
    for k,v := range map1 {
        if v == 1 {
            sum=sum+k
        }
    }
    return sum
    
}
