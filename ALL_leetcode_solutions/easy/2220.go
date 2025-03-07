// easy 2220 
// minimum bip flips to ...

// 87.5 acc rate
// JAN 13, 2025

// my alg: beats 100% cpu, beats only 4% RAM
// ok


import "fmt"

func abs (n int) int {
    ans := 0
    if n < 0 {
        ans = n * -1
    } else {
        ans = n + 0
    }
    
    return ans
}

func tenP (n int) int {

    ans:=1
    for i:=0; i<n; i++ {
        ans = ans * 10
    }

    return ans
}


func to_binary (n int) []int {

    ans := []int{}
    q := n
    remainder := 0

    for q > 0 {
      //fmt.Println(q % 2)
      remainder = q % 2
      q = q / 2
      ans = append(ans,remainder)
    }

    return ans

}

func minBitFlips(start int, goal int) int {

    //fmt.Println(to_binary(start))
    //fmt.Println(to_binary(goal))

    ans_p1 := 0
    ans_p2 := 0
    bin_start := to_binary(start) // binary list made from goal
    bin_goal := to_binary(goal) // binary list made from start
    L1 := len(bin_start)
    //fmt.Println(L1)
    L2 := len(bin_goal)
    //fmt.Println(L2)
    short_L := 0
    long_L := 0

    // see which is longer start or goal
    if L1 > L2 {
        short_L = L2
        long_L = L1
    } else {
        short_L = L1
        long_L = L2
    }

    // answer part 1 = number of 1s in the longer number in the digits that are not 
    // part of the common bits
    // e.g. 
    // start [0 0 1]
    // goal  [1 0 1 0 1]
    // only the 5th "1" in the goal counts
    if L2 > L1 { // goal is longer
       for i:=short_L ; i < long_L ; i++ {
          if bin_goal[i] == 1 {
            ans_p1++
          }
       }
    }
    if L1 > L2 { // start is longer
       for i:=short_L ; i < long_L ; i++ {
          if bin_start[i] == 1 {
            ans_p1++
          }
       }
    }
    //fmt.Println("ans_p1",ans_p1)

    
    // answer part 2
    // count how many bits are different among the common bits
    
    for i :=0 ; i < short_L ; i++ {
        if bin_start[i] != bin_goal[i] {
            ans_p2++
        }
    }
    //fmt.Println("ans_p2",ans_p2)
    
    ans := ans_p1 + ans_p2

    return ans
    
}
