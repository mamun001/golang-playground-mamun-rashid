// MEDIUM 1833 , ELO 1253
// maximum ice cream bars 
// my alg CPU > 20%, RAM > 98%!!

// did this in less than 5 minutes w/o even a syntax error

import "sort"

func maxIceCream(costs []int, coins int) int {
    sort.Ints(costs)

    at_hand := coins
    ans := 0

    for i:=0;i<len(costs) && at_hand > 0 ;i++ {
        if costs[i] <= at_hand {
            ans++
            at_hand = at_hand - costs[i]
        } else {
            break
        }
    }


    return ans
    
}
