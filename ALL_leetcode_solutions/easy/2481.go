// easy 2481, ELO 1246
// minimum cuts to divide a circle

// it can proven that 
// for even n, answer is n/2
// for odd n, answer n
// edge case 1, where the answer is 0

func numberOfCuts(n int) int {

    if n == 1 {
        return 0
    }

    if n % 2 == 1 {
        return n
    } else {
        return n/2
    }

    return -1
    
}
