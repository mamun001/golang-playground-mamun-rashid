// easy 2932 , ELO 1246
// Maximum Strong Pair XOR I
// Had to lookup how to get XOR of two integers

func abs(n int) int {
    if n < 0 {
        return n * -1
    }
    return n
}

func is_strong(n1,n2 int) bool {
    if abs(n1-n2) <= min(n1,n2) {
        return true
    }
    return false
}

func maximumStrongPairXor(nums []int) int {

    //fmt.Println(min(5,1))

    max:=0
    for i:=0;i<len(nums)-1;i++ {
        for j:=i;j<len(nums);j++ {
            //fmt.Println(i,j)
            if is_strong(nums[i],nums[j]) && nums[i]^nums[j] > max {
                max = nums[i]^nums[j]
            }
        }
    }


    return max
    
}
