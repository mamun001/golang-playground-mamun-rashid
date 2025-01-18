// easy 3248 snake in matrix
// ELO 1171
// JAN 17, 2025


func finalPositionOfSnake(n int, commands []string) int {

    // down = +n
    // up = -n
    // right = +1
    // left = -1

    ans := 0

    for _,v := range commands {
        if v == "UP" {
            ans = ans-n
        }
        if v == "DOWN" {
            ans = ans+n
        }
        if v == "RIGHT" {
            ans = ans+1
        }
        if v == "LEFT" {
            ans = ans-1
        }
    }

    return ans
    
}

