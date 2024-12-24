// easy 3190 find minimum operations to make al elements divisible by 3
// acc rate 89.1%

// DEC 24, 2024






func minimumOperations(nums []int) int {

    ans := 0

    // if the element is already divisble by zero, then no action needed
    // otherwise, in all cases, you either add 1 or subtract 1
    for _,v := range nums {
        if v % 3 != 0 {
            ans = ans + 1
        }

    }

    return ans
    
}
