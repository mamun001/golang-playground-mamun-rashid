// 72 pct acc rate

// easy level, 2595

// DEC 26, 2024

import "fmt"

func evenOddBit(n int) []int {

    ans := []int{}
    binary_representation := []int{}

    var remainder int
    remainder = 0
    var quotient int
    quotient = n // we start with n and keep dividing by 2 below

    for quotient != 0 {
        
        remainder = quotient % 2
        quotient = quotient / 2
        //fmt.Println(remainder)
        binary_representation = append(binary_representation, remainder)


    }

    // binary_represenattion above is actually reverse, so must reverse the slice

    // because says index is right to left, We don't need to reverse
    //reversed := []int{}
    //for i := len(binary_representation) -1 ; i > -1 ; i-- {
        //reversed = append(reversed, binary_representation[i])
    //}
    
    //fmt.Println(binary_representation)
    //fmt.Println(reversed)


    even := 0
    odd := 0

    for i,v := range binary_representation {
        //fmt.Println(i,v)
        if v == 1 {
            //fmt.Println("v is 1")
            if i % 2 == 0  {
                //fmt.Println("even")
                even = even + 1
            } else {
                //fmt.Println("odd")
                odd = odd + 1
            }
        }
    }


   ans = append (ans,even)
   ans = append (ans,odd)

   return ans
}
