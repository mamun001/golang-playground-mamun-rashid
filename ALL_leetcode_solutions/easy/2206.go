/*
easy 2206, ELO 1223
divide array into equal pairs

*/


func divideArray(nums []int) bool {

    //ans:= true

    map1:= make(map[int]int)

    for i:=0;i<len(nums);i++ {
        map1[nums[i]]++
    }

    for _,v := range map1 {
        if v % 2 == 1 {
            return false
        }
    }

    return true
    
}
