// easy 1979, ELO 1184
// find greatest common divisor

// jan 23, 2025



func findGCD(nums []int) int {

    min := 1001
    max := 0
    gcd := 1

    for i:=0 ; i <len(nums); i++ {
        if nums[i] > max {
            max = nums[i]
        }

        if nums[i] < min {
            min = nums[i]
        }
    }

    //fmt.Println(min,max)

    for i:=1 ; i <= max ; i++ {
        if min % i == 0 && max % i == 0 {
            gcd = i
        }
    }

    return gcd
    
}

