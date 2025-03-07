// MEDIUM , 15, 3SUM, Part of Neetcode 250
// Needed help from youtube
// even with part of the code
// SO, this one needs repeating 10 times

import "sort"

func threeSum(nums []int) [][]int {

    sort.Ints(nums)
    //fmt.Println(nums)

    ans :=[][]int{}
    


    for i:=0;i<len(nums);i++ {
        //fmt.Println(" ")
        //fmt.Println("i",i)
        if i>0 && nums[i] == nums[i-1] {
            continue
        }
        l:=i+1
        //fmt.Println("l",l)
        r:=len(nums)-1
        //fmt.Println("r",r)
        for l<r {
          sum := nums[i] + nums[l] + nums[r]
          if sum == 0 {
            //fmt.Println(nums[i],nums[l],nums[r])
            temp := []int{nums[i],nums[l],nums[r]}
            ans = append(ans,temp)
            l++
            //fmt.Println("l",l)
            for nums[l] == nums[l-1] && l<r {
                l++
                 //fmt.Println("l",l)
            }
          } else if sum > 0 {
            r--
            //fmt.Println("r",r)
          } else if sum < 0 {
            l++
             //fmt.Println("l",l)
            
          }
        }
    }

    

    return ans
    
}
