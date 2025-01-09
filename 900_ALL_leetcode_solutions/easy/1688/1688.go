// easy 1688
// Count of Matches Tournament
// 85.8 acc rate
// JAN 9, 2025



func numberOfMatches(n int) int {

    // Use of recursion makes sense here
    // formulas are literally given in the description of the problem

    ans := 0

    if n == 1 {
        return 0
    } else if n == 2 {
        return 1
    } else if n % 2 == 0 { // even
       return n / 2 + numberOfMatches (n/2)
    } else if n % 2 == 1 { // odd
         return (n -1) / 2 + numberOfMatches ((n-1)/2 + 1)
    }



    return ans
    
}

