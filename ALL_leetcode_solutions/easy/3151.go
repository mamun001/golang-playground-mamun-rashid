// easy 3151 special array I
// ELO 1154
// JAN 17, 2025

func isArraySpecial(nums []int) bool {

    ans := true

    L := len(nums)

    last := 1 // 1 indicates odd, 0 indicates even
    this := 0 // to start off


    // set the first "last" value
    if nums[0] % 2 == 1 {
        last = 1
    } else {
        last = 0
    }

    // by problem statement
    if L == 1 {
        ans = true
    }


    // My alg: each item must alternate between odd and even for the array to be special
    for i:= 1 ; i < L ; i++ {
        this = nums[i] % 2 
        if this == last {
            return false
        } else {
            last = this
        }   
    }
    return ans
    
}

