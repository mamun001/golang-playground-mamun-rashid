// 2610 , MEDIUM , ELO 1373
// convert an array into a 2D array with conditions

func is_empty(map1 map[int]int) bool {
    for _,v := range map1 {
        if v > 0 {
            return false
        }
    }
    return true
}

func findMatrix(nums []int) [][]int {

    map_num_freq := make(map[int]int)

    //fmt.Println(is_empty(map_num_freq))

    for i:=0;i<len(nums);i++ {
        map_num_freq[nums[i]]++
    }
    
    //fmt.Println(map_num_freq)
    //fmt.Println(is_empty(map_num_freq))

    ans:=[][]int{}

    // k is 1,2,3 etc
    // until every single value in map is zero, we keep taking one from each key 
    // and form a row. (and decrement value by 1)

    for is_empty(map_num_freq) == false {
      row := []int{}
      for k,v := range map_num_freq {
          //row := []int{}
          if v > 0 {
              row = append(row,k)
              map_num_freq[k]--
          }
          
      }
      // fmt.Println("row",row)
      ans = append(ans,row)
      //fmt.Println(map_num_freq)
    }

    return ans
    
}
