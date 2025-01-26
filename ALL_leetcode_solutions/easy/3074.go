// easy 3074, ELO 1197
// apple redistribution into

// jan 24, 2025

// get sum, sort capacity, pick the largest boxes



import "sort"

func minimumBoxes(apple []int, capacity []int) int {

    ans := 0

    apples := 0
    filled := 0
    boxes := 0

    for i:=0; i<len(apple); i++ {
        apples = apples + apple[i] // total number of apples
    }

    //fmt.Println(apples)

    sort.Ints(capacity)

    for i:=len(capacity)-1; i>-1; i-- {
        boxes++
        filled = filled + capacity[i]
        //fmt.Println(i,capacity[i],filled)
        if filled >= apples {
            return boxes
        }
    }


    return ans
    
}

