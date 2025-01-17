// easy 3024 type of triangle
// ELO 1134

// JAN 16, 2025

import "sort"

func triangleType(nums []int) string {

    sort.Ints(nums) // sorting makes rest of the code easier and easier to read

    ans := ""

    a := nums[0] // smallest
    b := nums[1] // medium
    c := nums[2] // largest

    // can it form a triangle?
    if a+b <=c || b+c <=a || c+a <=b {
        return "none"
    }

    if a == b && b == c {
        return "equilateral"
    }

    if a != b && b != c && c != a {
        return "scalene" // all unequal
    }

    if a == b && b != c { // we can do this a,b,c are sorted
        return "isosceles"
    }


    if a !=b && b == c { // we can do this a,b,c are sorted
        return "isosceles"
    }

    return ans
    
}
