// easy 2037 minimum number of moves
// 87.4 acc rate

// JAN 15, 2025

// my fancy alg: beats 100% cpu , BUT beats only 11% RAM


import "fmt"
import "sort"

func abs (n int) int {
    ans := n
    if n < 0 {
        ans = n * -1
    }
    return ans
}


func minMovesToSeat(seats []int, students []int) int {

    ans := 0

    sort.Ints(seats)
    sort.Ints(students)

    // just add up the sortest distance to closet seat (kind of)

    for i:=0 ; i < len(seats) ; i++ {
        ans = ans + abs( students[i] - seats[i])
    }



    return ans
    
}

