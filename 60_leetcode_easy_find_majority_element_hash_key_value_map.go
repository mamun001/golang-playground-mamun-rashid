
func majorityElement(nums []int) int {
    
    answer := 0
    maxvalue := 0
    map1 := make(map[int]int)   
    
    // fill up key,value with data from nums
    
    for i:=0;i<len(nums);i++ {
       x:=nums[i]
       if _, found := map1[x]; found { 
          map1[x]=map1[x] + 1
       } else {
           map1[x]=1
       }
     }
    
    // now find value that is the max, return its key
    
    for key, value := range map1 {
        if value > maxvalue {
            maxvalue = value
            answer = key
        }
    }
    
    return answer
    
}
