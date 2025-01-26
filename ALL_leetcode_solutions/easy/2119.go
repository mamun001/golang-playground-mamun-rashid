// easy 2119 ELO 1187
// a number aftre a double reversal

// clever math alg : cpu beats 100%, RAM beats only 17%

func isSameAfterReversals(num int) bool {

    ans := true

    if num == 0 {
        return true
    }

    if num % 10 == 0 {
        return false
    } else {
        return true
    }

    return ans
    
}
