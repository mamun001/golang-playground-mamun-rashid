// easy 2848, ELO 1230
// Points That Intersect With Cars

func numberOfPoints(nums [][]int) int {

    map1 := make(map[int]int) // 0 means no car, 1 means car

    m := len(nums)
    for i:=0;i<m;i++{
        for j:=nums[i][0];j<=nums[i][1];j++{
            map1[j] = 1
        }
    }

    return len(map1)
}

