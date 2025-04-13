// easy 2180 , ELO 1257
// count integers with even digit sum

func digit_sum(n int) int {
    sum:=0
    q := n
    for q > 0 {
        sum = sum + q % 10
        q = q / 10
    }
    return sum
}

func countEven(num int) int {

    //fmt.Println(digit_sum(123))
    
    ans:=0

    for i:=1;i<=num;i++{
        if digit_sum(i) % 2 == 0 {
            ans++
        }
    }
    return ans
    
}
