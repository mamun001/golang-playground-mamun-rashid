// easy 3099
// ELO 1084 (lowest)

// JAN 15, 2025


func sumOfTheDigitsOfHarshadNumber(x int) int {

    ans := -1

    if x == 100 {
        ans = 1
    } else if x < 10 {
        ans = x
    } else {
        d1 := x / 10
        d2 := x % 10
        if x % (d1+d2) == 0 {
            ans = d1+d2
              }
    }
    return ans
}

