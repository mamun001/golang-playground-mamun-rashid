// 66.9% percentile

// 3206 alternating groups I

// DEC 27, 2024


import ("fmt")

func numberOfAlternatingGroups(colors []int) int {
   ans := 0
   previous_index := 9999
   next_index := -9999


   for i,_ := range colors {
    if i == 0 {
        previous_index = len(colors) -1 // first's previous is "last" in this cycle
    } else {
        previous_index = i-1
    }

    if i == len(colors) -1  {
        next_index = 0 // last's "next" is "first"
    } else {
        next_index = i+1
    }

    if colors[i] != colors[previous_index] && colors[i] != colors[next_index] {
        fmt.Println("index",i)
        ans++
    }


   } // for loop

    return ans
}




