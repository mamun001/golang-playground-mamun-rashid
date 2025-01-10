// easy 2520
// count the digits that divide a number
// 85% acc rate

// JAN 9, 2025

// My algorith is VERY inefficent

// I did not relaize I could modify num ; also num % 10 gives the digit, not all the complex stuff I did

import "fmt"



func countDigits(num int) int {

    ans := 0

    ten_exp := 1
    digit_count := 0
    this_digit := 10
    remainder := num


    // count the number of digits in num
    quo := num
    for i:=0 ; quo > 0 ; i++ {
        quo = quo / 10
        //fmt.Println(quo)
        digit_count++
    }
    //fmt.Println(digit_count)


    
    for i:=digit_count ; this_digit > 0 ; i-- {
        fmt.Println("digit number from the right",i)
        for j:=1; j <i ; j++ {
           ten_exp = ten_exp * 10
        } // for inner
        fmt.Println("will divide this by this", remainder, ten_exp)
        this_digit = remainder / ten_exp
        
        fmt.Println("this_digit", this_digit)
        if this_digit == 0 {
            fmt.Println("breaking")
            break
        }
        if num % this_digit== 0 {
            ans++
        }
        remainder = remainder % ten_exp
        fmt.Println("remainder", remainder)
        ten_exp = 1
    } // for outer
    
    

    return ans
    
}
