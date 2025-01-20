// MEDIUM 985 ELO 1198
// Sum of Even Numbers After Queries
// JAN 19, 2025
// MY ALG , CPU beats 0% (worst, 6 ms) , RAM beats 100%


import "fmt"

func sum_of_evens(arr []int) int {
    ans :=0
    for i:=0; i<len(arr);i++ {
       if arr[i] % 2 == 0 {
        ans = ans + arr[i]
       }
    }

    return ans
}

func sumEvenAfterQueries(nums []int, queries [][]int) []int {

    ans := []int{}

    for i:=0; i < len(queries) ; i++ {
        
        how_much := queries[i][0]
        where := queries[i][1]
        //fmt.Println(where,how_much)
        nums[where] = nums[where] + how_much
        //fmt.Println(nums)
        this_sum := sum_of_evens(nums)
        //fmt.Println(this_sum)
        ans = append(ans,this_sum)
    }

    return ans
    
}
