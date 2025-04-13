// easy 3379, ELO 1257
// transformed array

func constructTransformedArray(nums []int) []int {

    result:=[]int{}

    for i:=0;i<len(nums);i++ {
        result=append(result,0)
    }

    for i:=0;i<len(nums);i++ {
        index:=(i+nums[i])
        for index < 0 { // sometimes the value is outrageous like -10, so its loop until
            index=index+len(nums)
        }
        for index > len(nums)-1 { // sometimes the value is outrageous like -10, so its loop until
            index=index-len(nums)
        }
        //fmt.Println(index)
        result[i]=nums[index]
    }


    return result
    
}
