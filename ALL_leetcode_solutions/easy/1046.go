// easy 1046, ELO 1172, Last Stone Weight
// Jan 20, 2025
// my alg is not very efficient because I am sorting N times!
// cpu = beats 4%, RAM = beats 10%


import "sort"

func lastStoneWeight(stones []int) int {

    N := len(stones)
    sort.Ints(stones)
    fmt.Println(stones)
    y := 0
    x := 0
    all_done := false

    if N == 1 {
        return stones[0]
    }

    for i:=0 ; all_done == false ; i++ {
        y = stones[N-1]
        x = stones[N-2]
        if x == y {
            stones[N-1] = 0 // y
            stones[N-2] = 0 // x
        } else if x != y {
            stones[N-2] = 0 // x
            stones[N-1] = y - x // y
        }
        fmt.Println(stones)
        sort.Ints(stones)
        if stones[N-2] == 0 { // at most one non zero stone lest
           all_done = true
        }
    } // for

    return stones[N-1]
    
}

