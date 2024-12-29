// 6th easiest per chatgpt

// 53.3 pct acc rate

// Problem number 70 , easy

// DEC 28, 2024

// This is just fibonacci series

// However my recursive func ONLY worked until n =43
// at 44 , it ran out of time
// I had to use Dynamic Programming
// Code for that is basically a code in Python FROM YOUTUBE that I converted to Golang
// SO, I need to try this again in a month

func climbStairs(n int) int {

    ONE := 1
    TWO := 1

    for i := 0 ; i < n-1 ; i ++ {
        // ONE gets larger
        // TWO gets ONE's value
        // so really here, TWO is "behind" ONE

        // for example n=1 , answer 1
        //             n=2 , answer 2 
        //             n=3 , answer 3 (here, ONE is 2 and TWO 1) and we return TWO+ONE = 2 +1 = 3
        temp := ONE
        ONE = ONE + TWO
        TWO = temp 

    }

    return ONE
        
    
}
