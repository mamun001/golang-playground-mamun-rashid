// easy 2455
// average value of even
// ELO 1151

// JAN 17, 2025


import "fmt"
func averageValue(nums []int) int {

    ans := 0

    L := len(nums)

    sum := 0
    count := 0

    for i:=0 ; i<L ; i++ {
        fmt.Println(nums[i])
        if nums[i] % 3 == 0 && nums[i] % 2 == 0 {
            fmt.Println(i)
            sum = sum + nums[i]
            count++
        }

    }

    if count == 0 {
        return 0
    } else {
       ans = sum / count
    }

    return ans
    
}

