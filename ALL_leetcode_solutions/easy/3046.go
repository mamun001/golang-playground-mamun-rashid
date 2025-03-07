// easy 3046, ELO 1212
// Split the array

func isPossibleToSplit(nums []int) bool {

    map1 := make(map[int]int)

    for i:=0;i<len(nums);i++ {
        map1[nums[i]]++
    }

    //fmt.Println(map1)
    for _,v := range map1 {
        if v > 2 {
            return false
        }
    }



    return true
    
}

