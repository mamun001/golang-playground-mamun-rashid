/*
easy, 2535, ELO 1222
difference between element sum and


*/

func abs (n int) int {
    if n < 0 {
        return n * (-1)
    }

    return n
}

func digit_sum(n int) int {
    ans:= 0
    q := n
    d := 0


    for q>0 {
        d = q % 10
        ans = ans + d
        q = q/10
    }

    return ans
}

func differenceOfSum(nums []int) int {

    
    e_sum:=0
    d_sum:=0

    //fmt.Println(digit_sum(0))
    for i:=0; i<len(nums);i++ {
        e_sum = e_sum + nums[i]
        d_sum = d_sum + digit_sum(nums[i])
    }


    return abs(e_sum - d_sum)
    
}
