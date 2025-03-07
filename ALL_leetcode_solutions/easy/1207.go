// 77 percent acceptance rate

// 1207 Unique Number of Occurrences

// again, the problem is no good

// DEC 25, 2024

import "fmt"

func in_slice (s []int, new_item int) bool {
    ans := false

    for _,v := range s {
        if v == new_item {
            ans = true
        }
    }

    return ans
}

func uniqueOccurrences(arr []int) bool {

    keeping_track_freq :=make(map[int]int)
    //keeping_track_freq[1] = 2

    // lets add the frequencies of values in arr
    for _,v := range arr {
        keeping_track_freq[v]++
    }

    //fmt.Println(keeping_track_freq)

    // IF there is ANY repeats in "values" of keeping_track , then the answer is TRUE
    // if ALL values are uniq (no repeats) , the the answer is FALSE

    // REPEAT YES? = FALSE
    // REPEAT NO? = TRUE

    // we start with "true" (no repeat), then we see our first repeat , we set to "false" and we break

    ans := true
    keeping_track_uniq := []int{} // this is essentially a set, because we break if we see that item already exists in slice

    for _,v := range keeping_track_freq {
        if in_slice(keeping_track_uniq,v) {
            ans = false // we found our first repeat
            break
        } else {
            keeping_track_uniq = append(keeping_track_uniq, v) // still no repeat, add v to SET
        }
    }

    fmt.Println(keeping_track_uniq)
    
    return ans
}


