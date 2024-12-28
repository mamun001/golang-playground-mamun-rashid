// 64.7 pct

// easy, 2833

// Furthest Point from origin


import "fmt"

func furthestDistanceFromOrigin(moves string) int {

    ans := 0
    L := 0
    R := 0
    ANY := 0

    for _,v := range moves {
        if v == 'L' {
            L++
        } else if v == 'R' {
            R++
        } else {
            ANY++
        }

    }

    // came up with this algorith by myself
    if L >=R {
        ans = L + ANY - R
    } else {
        ans = R + ANY - L
    }

    fmt.Println(L,R,ANY)
    return ans
    
}
