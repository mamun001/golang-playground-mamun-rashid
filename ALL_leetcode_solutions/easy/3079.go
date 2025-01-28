// easy, 3079, find the sum of encrypted integers
// ELO 1190

// JAN 25, 2025

// amy alg very poor: CPU > 21%, RAM > 9%

import "strconv"

func ld(n int) int {
    // return largest digit in n
    s := strconv.Itoa(n)
    max :=0
    digit :=0
    for i,_ := range s {
        digit,_ = strconv.Atoi(string(s[i]))
        if digit > max {
            max = digit
        }
    }
    return max
}

func replace_all(n int, z int) int {
    // replace all digits with "i"
    s := strconv.Itoa(n)
    L := len(s)
    new_s := ""
    char := ""
    for i:=0;i<L;i++ {
        //fmt.Println(i,string(i))
        char = strconv.Itoa(z)
        new_s = new_s + char
    }
    //fmt.Println(new_s)
    ans,_ := strconv.Atoi(new_s)
    return ans
}

func sumOfEncryptedInt(nums []int) int {

    ans :=0
    //fmt.Println(ld(987))
    //fmt.Println(" ")
    //fmt.Println(replace_all(123,3))
    //fmt.Println(replace_all(986544,9))
    //fmt.Println(replace_all(1,ld(1)))
    
    for i:=0;i<len(nums);i++{
        nums[i] = replace_all(nums[i],ld(nums[i]))
    }

    //fmt.Println(nums)

    for i:=0;i<len(nums);i++{
        ans = ans + nums[i]
    }

    return ans
    
}
