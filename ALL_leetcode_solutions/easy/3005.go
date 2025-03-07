// easy 3005, Count Elements With Maximum Frequency
// ELO 1217

func maxFrequencyElements(nums []int) int {

   map_freq := make(map[int]int)

   for i:=0;i<len(nums);i++ {
     map_freq[nums[i]]++
   }

   max := 0
   for _,v := range map_freq {
    if v > max {
        max = v
    }
   }

   ans := 0
   for _,v := range map_freq {
    if v == max {
        ans = ans + v
    }
   }


    return ans
    
}
