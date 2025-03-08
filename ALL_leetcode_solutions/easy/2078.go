// easy 2078 , ELO 1241
// Two furthest houses with different colors
// brute force alg works b/c of the limits guven

func abs (n int) int {
    if n < 0 {
        return n* (-1)
    }
    return n
}

func maxDistance(colors []int) int {

    n:=len(colors)
    max:=0
    

    for i:=0;i<n;i++ {
        for j:=0;j<n;j++ {
            if colors[i] != colors[j] && abs(j-i) > max {
               max = abs(j-i)
            }
        }
    }


    return max
    
}
