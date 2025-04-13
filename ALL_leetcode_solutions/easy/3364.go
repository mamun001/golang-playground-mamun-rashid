// easy 3364 , ELO 1301
// minimum positive sum subarray

func sum(ar []int) int {
    ans:=0
    for i:=0;i<len(ar);i++{
        ans=ans+ar[i]
    }
    return ans
}

func minimumSumSubarray(nums []int, l int, r int) int {


min:=99999999
  for z:=l;z<=r;z++ {
    for i:=0;i+z<=len(nums);i++{
        temp:=nums[i:i+z]
        //fmt.Println(sum(temp))
        if sum(temp) > 0 && sum(temp) < min {
            min = sum(temp)
        }
    }
  }

if min == 99999999 {
    return -1
}

return min
    
}
