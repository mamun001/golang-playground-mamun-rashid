

// leetcode easy 1431
// for loop

func kidsWithCandies(candies []int, extraCandies int) []bool {
    a := make([]bool, len(candies))
    var max int = 0
    
    for i:=0; i < len(candies); i++ {
        if candies[i] > max {
            max = candies[i]
        }
    }
    
    
    
    for i:=0; i < len(candies); i++ {
        if (candies[i] + extraCandies) > max-1 {
        a[i] = true
    }
    }
    
    return a
}
