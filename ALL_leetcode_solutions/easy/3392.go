// easy 3392 ELO 1201
// 3392 count subarrays of length

func countSubarrays(nums []int) int {



   ans := 0

   for i:=0;i<len(nums)-2; i++ {
      //fmt.Println(nums[i+2])
      if (nums[i] + nums[i+2]) * 2 == nums[i+1] {
        fmt.Println(nums[i],nums[i+1],nums[i+2]) // doing the other way with /2 causes edge case problems with rounding
        ans++
        fmt.Println(ans)
      }

   }


    return ans
    
}
