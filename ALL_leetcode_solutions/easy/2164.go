// easy 2164 , ELO 1253
// Sort Even and Odd Indices independently

import "sort"


func rev_arr(ar []int) []int {

    ans:=[]int{}

    for i:=len(ar)-1;i>-1;i--{
        ans=append(ans,ar[i])
    }

    return ans
}


func sortEvenOdd(nums []int) []int {

    odds := []int{}
    evens := []int{}

    for i:=0;i<len(nums);i++ {
        if i % 2 == 1 {
            odds = append(odds,nums[i])
        }
        if i % 2 == 0 {
            evens = append(evens,nums[i])
        }
    }

    sort.Ints(odds)
    rev_odds:=rev_arr(odds)
    sort.Ints(evens)

    //fmt.Println(rev_odds)
    //fmt.Println(evens)


    ans:=[]int{}

    for i:=0;i<len(nums);i++{
        if i % 2 == 0 {
           ans=append(ans,evens[i/2])
        }
        if i % 2 == 1 {
           ans=append(ans,rev_odds[i/2])
        }
    }

    

    return ans
    
}
