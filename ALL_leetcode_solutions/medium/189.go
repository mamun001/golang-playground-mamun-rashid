// MEDIUM 189, no ELO
// part of neetcode 250
// rotate array
// done in 12-15 minutes

func rotate(nums []int, k int)  {

    temp:=[]int{}
    n := len(nums)
    real_k := k % n

    for i:=0;i<real_k;i++ {
        temp = append(temp,nums[n-real_k+i])
    }
    //fmt.Println(temp)

    for i:=n-1;i >= real_k;i-- {
        nums[i] = nums[i-real_k]
    }

    for i:=0;i<real_k;i++ {
        nums[i] = temp[i]
    }


    
}
