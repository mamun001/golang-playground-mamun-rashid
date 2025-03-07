// easy , 2859, ELO 1218
// Sum of Values at Indices With K set bits

import "strconv"

func binary(n int) string {

    if n == 0 {
        return "0"
    }

    if n == 1 {
        return "1"
    }

    ans := ""
    //fmt.Println("start")
    q := n / 2
    r := n % 2
    ans = strconv.Itoa(r) + ans
    //fmt.Println(r)

    for q > 0 {
        r = q % 2 
        q = q / 2
        //fmt.Println(r)
        ans = strconv.Itoa(r) + ans
    }

    //fmt.Println("finish")
    return ans
}

func set_bits (n int) int {
    ans := 0
    str := binary(n)
    for i:=0; i<len(str);i++ {
        if str[i] == '1' {
            ans++
        }
    }
    return ans
}

func sumIndicesWithKSetBits(nums []int, k int) int {

    //n := 33
    //fmt.Println(binary(35))
    //fmt.Println(set_bits(n))

    sum := 0
    for i:=0; i<len(nums);i++ {
        if set_bits(i) == k {
            sum = sum + nums[i]
        }
    }

    return sum
    
}
