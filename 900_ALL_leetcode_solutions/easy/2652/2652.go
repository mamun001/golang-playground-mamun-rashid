// easy 2652
// 85.3 acc rate
// sum multiples



func sumOfMultiples(n int) int {

    ans := 0

    for i:= 1 ; i <= n ; i++ {
        if i % 3 == 0 || i % 5 == 0 || i % 7 == 0 {
            ans = ans + i
        }
    }

    return ans
    
}
