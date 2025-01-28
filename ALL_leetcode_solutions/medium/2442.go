// MEDIUM 2442 Count Number of Distint Integers After Reverser Operations
// my first alg failed because I did not use map for lookup and it took too long
// Once I changed to map AND did the lookup while I was reversing ("one" run through array)
// performance increased

// ELO = 1218

// alg perf: beats CPU 10% (231 ms) Close!
//           beats RAM 71%

// Jan 27, 2025

import "strconv"


// func to reverse an integer
func rev (n int) int {
    s1 := strconv.Itoa(n)
    // reverse the string so we can reverse the int effectively
    s_rev := ""
    for i:=len(s1)-1; i > -1 ; i-- {
        s_rev = s_rev + string(s1[i])
    }
    //fmt.Println(s_rev) // 87654321
    reversed_int,_ := strconv.Atoi(s_rev)
    //fmt.Println(reversed_int) // 87654321

    return reversed_int
}


func countDistinctIntegers(nums []int) int {

    // map of unique integers thus far
    // my first attempt failed beacuse I did not use map
    map_uniq := make(map[int] bool)

   // populate map as we reverse, in "one" shot
   for i:=0; i<len(nums) ; i++ {
     //fmt.Println(rev(nums[i]))
     if map_uniq[nums[i]] != true  {
        //fmt.Println(nums[i],false)
        map_uniq[nums[i]]  = true
     }
     if map_uniq[rev(nums[i])] != true  {
        //fmt.Println(nums[i],false)
        map_uniq[rev(nums[i])]  = true
     }
    }

    return len(map_uniq)
    
}

