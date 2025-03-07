// easy 2980 ELO 1234
// check if bitwise OR has trailing zeros

import "strconv"

func bin(n int) string {
    ans := ""

    if n == 0 {
        return "0"
    }

    if n == 1 {
        return "1"
    }

    q := n
    r := 0

    for q > 0 {
        r = q % 2 
        c := strconv.Itoa(r)
        ans = c + ans
        q = q / 2
    }

    return ans
}


func hasTrailingZeros(nums []int) bool {


    //fmt.Println(nums[0] | nums[1])
    //fmt.Println(bin(32))

    for i:=0;i<len(nums);i++ {
        for j:=0;j<len(nums);j++ {
            if i != j {
                or_res := nums[i] | nums [j]
                bin_or_res := bin(or_res)
                if string(bin_or_res[len(bin_or_res)-1]) == "0" {
                    return true
                }
            }
        }
    }

    return false
    
}
