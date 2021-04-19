
// easy LC
// implements set
// https://leetcode.com/problems/single-number/

func is_in_set(s []int, item int) bool {
    answer := false
    
    for i:=0 ; i < len(s); i++ {
      if item == s[i] {
          answer = true
      }
    }
    
    return answer
}

func if_in_then_remove(s []int, item int) []int{
    answer := s
    
    for i:=0 ; i < len(s); i++ {
      if item == s[i] {
          s[i] = s[len(s)-1]
          answer= s[:len(s)-1]
      }
    }
    
    return answer
}




func singleNumber(nums []int) int {
    
    //answer := 0
    
    
    set_of_ones:= make([]int, 0)
    //set_of_ones = append(set_of_ones, 5)
    //fmt.Println(is_in_set(set_of_ones, 5))
    //fmt.Println()
    //set_of_ones=if_in_then_remove(set_of_ones, 5)
    //fmt.Println(is_in_set(set_of_ones, 5))
    
    for i:=0; i < len(nums); i++ {
        if is_in_set(set_of_ones,nums[i]) == true {
            set_of_ones=if_in_then_remove(set_of_ones, nums[i])
        } else {
            set_of_ones = append(set_of_ones, nums[i]) 
        }   
    }
    
    
    return set_of_ones[0]
    
}
