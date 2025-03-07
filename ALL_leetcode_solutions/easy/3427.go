// easy 3427, ELO 1215
// sum of vaiable length subarrays

func start(ar []int, i int) int {

    ans := 0
    iminusnumi := i - ar[i]
    if iminusnumi > 0 {
        return iminusnumi
    } else {
        return 0
    }
    return ans
}

func subarraySum(nums []int) int {

    //fmt.Println(start(nums,0))
    //fmt.Println(start(nums,1))
    //fmt.Println(start(nums,2))

    sum :=0
    for i:=0;i<len(nums);i++ {
        st := start(nums,i)
        for j:=st;j<=i;j++ {
           sum = sum + nums[j]
        }
        //fmt.Println(sum)
    }
    

    return sum
    
}

