

func two_to_the_power (n int) int {
    // 2^n
    if n == 0 {
        return 1
    } else {
        return 2 * two_to_the_power (n-1)
    }
}
