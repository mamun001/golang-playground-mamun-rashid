// MEDIUM 2279 maximum bags with full
// My alrg: CPU > 42%
// RAM > 47% , so pretty good

// Jan 28, 2025

import "sort"

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {

    n := len(rocks)

    already_full := 0
    cap_left :=[]int{}

    for i:=0 ; i<n ; i++ {
        if rocks[i] == capacity[i] {
            already_full++
        } else {
            cap_left = append(cap_left,capacity[i]-rocks[i])
        }
    }
    //fmt.Println("already_full",already_full)
    //fmt.Println("cap_left",cap_left)

    sort.Ints(cap_left)
    rocks_left_to_pour := additionalRocks
    poured :=0
    new_full :=0
    //fmt.Println("sorted",cap_left)
    m := len(cap_left)
    //fmt.Println("len of cap_left m",m)
    //fmt.Println(" ")


    for i:=0; rocks_left_to_pour > 0 && i<m; i++ {
        //fmt.Println(" ")
        //fmt.Println("i",i)
        //fmt.Println("left_to_pour_count & how_many_this_time",rocks_left_to_pour, cap_left[i])
        if rocks_left_to_pour >= cap_left[i] { // if we are NOT "filling", then last one should not count
            new_full++
        }
        rocks_left_to_pour = rocks_left_to_pour - cap_left[i]
        poured = poured + cap_left[i]
        //fmt.Println("poured_rocks_count new_full_count", poured, new_full)
    }


    return already_full + new_full
    
}

