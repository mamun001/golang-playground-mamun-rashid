// leetcode 3285 find indices of stable mountains

// acc rate 86.9%

// DEC 24, 2024


func stableMountains(height []int, threshold int) []int {

    ans := []int{}

    for i,_ := range height {
        if i == 0 {
            continue // first item cannot be compared with previous because there is NO previous
        }
        if height[i-1] > threshold {
            ans = append(ans,i)
        }
    }
    
    return ans
}
