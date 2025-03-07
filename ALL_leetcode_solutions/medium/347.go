// MEDIUM , 347, no ELO
// part of Neetcode 250

// alg:
// gather frequencies in a map (map1)
// make an array of just frequencies
// sort it
// keep only the top K
// figure out loest frequency that would make the cut
// run through the map (map1) again and only return the indices that has frequeencies  >= least one that made the cut

import "sort"

func topKFrequent(nums []int, k int) []int {

    map1 := make(map[int]int)
    ans := []int{}

    for i:=0;i<len(nums);i++ {
        map1[nums[i]]++
    }

    //fmt.Println(map1)

    ar:=[]int{}

    for _,v := range map1 {
        ar = append(ar,v)
    }

    sort.Ints(ar)
    //fmt.Println("sorted",ar)

    ar = ar[len(ar)-k:]
    //fmt.Println("only the highest frequencies we need",ar)
    lowest := ar[0]
    //fmt.Println("lowest that makes the team",lowest)

    for i,v := range map1 {
        if v >= lowest {
            ans = append(ans,i)
        }
    }



   return ans

    
}
