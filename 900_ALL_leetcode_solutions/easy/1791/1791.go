// easy 1791
// find center of star graph
// my alg beats 100% cpu and 96% RAM!!
// common number in the 2 elements is all you need

// this is lazy coding, but it is such easy problem that it is ok

import "fmt"

func is_in_slice (s1 []int, n int) bool {
    ans := false
    for i:=0; i < len(s1); i++ {
        if s1[i] == n {
        ans = true
        }
    }

    return ans
}

func findCenter(edges [][]int) int {


    ans := 0

    slice1 := []int{}

    one := edges[0][0]
    two := edges[0][1]
    three := edges[1][0]
    four := edges[1][1]

    if is_in_slice (slice1, one) {
        ans = one
    } else {
        slice1 = append (slice1,one)
    }

    if is_in_slice (slice1, two) {
        ans = two
    } else {
        slice1 = append (slice1,two)
    }

    if is_in_slice (slice1, three) {
        ans = three
    } else {
        slice1 = append (slice1,three)
    }

    if is_in_slice (slice1, four) {
        ans = four
    } else {
        slice1 = append (slice1,four)
    }

    fmt.Println(edges[0][0])



    return ans
    
}

