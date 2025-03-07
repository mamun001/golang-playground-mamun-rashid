// easy, 2656, ELO 1208
// needed chatgpt help in using DP to make fib func not get stackoverflow using recursion
// so, this code does not use recursion

import "sort"

func fib(n int) int {

    if n <= 1 {
        return n
    }
    
    fib_slice := make([]int, n+1)
    fib_slice[0] = 0
    fib_slice[1] = 1

    for i:=2; i<=n;i++ {
        fib_slice[i] = i + fib_slice[i-1] 
    }

    return fib_slice[n]
}

func maximizeSum(nums []int, k int) int {

    sort.Ints(nums)

    big :=nums[len(nums)-1]

    ans:= big*k + fib(k-1)

    //fmt.Println(fib(10))

    return ans
    
}
