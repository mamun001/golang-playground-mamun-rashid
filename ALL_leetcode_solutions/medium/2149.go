// MEDIUM 2149, ELO 1235
// Rearrange Array Elemnts by sign


// my alg: two queues
// my alg: CPU > 6% 574 ms (close!)
// RAM beats 14%


func rearrangeArray(nums []int) []int {

    ans := []int{}

    pos := []int{}
    neg := []int{}

    for i:=0; i<len(nums); i++ {
        if nums[i] > 0 {
            pos = append(pos,nums[i])
        }
        if nums[i] < 0 {
            neg = append(neg,nums[i])
        }
    }

    //fmt.Println(pos)
    //fmt.Println(neg)

    next := 0
    for i:=0; i<len(nums); i++ {
      if next == 0 {
        ans = append(ans,pos[0])
        pos = pos[1:len(pos)]
        //fmt.Println(pos)
      }
      if next == 1 { // neg
        fmt.Println("at neg")
        ans = append(ans,neg[0])
        neg = neg[1:len(neg)]
        //fmt.Println(ans)
        //fmt.Println(neg)
      }

      if next == 0 {
        next = 1
        //fmt.Println("next")
      } else if next == 1 {
        next = 0
      }
    }

    


    return ans
    
}


