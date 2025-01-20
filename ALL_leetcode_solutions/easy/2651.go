// easy 2651, calculate delayed arrival
// ELO 1168

// JAN 18, 2025
// I can't believe this is on leetcode
// I am legend

func findDelayedArrivalTime(arrivalTime int, delayedTime int) int {

    ans := arrivalTime + delayedTime
    if ans > 23 {
        ans = ans -24
    }
    
    return ans
}
