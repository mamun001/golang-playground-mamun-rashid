// easy 1486, ELO 1180
// XOR Operation in an array

// JAN 23, 2025

func xorOperation(n int, start int) int {

    ans := 0

    nums := make([]int,n)

    for i:=0; i<n; i++ {
        nums[i] = start + 2 * i
    }

    for i:=0; i<n; i++ {
        ans = ans ^ nums[i]
    }
    //a := 2 ^ 4 ^ 6
    //fmt.Println(nums)

    return ans
    
}
