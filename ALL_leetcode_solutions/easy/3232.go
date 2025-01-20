// ELO 1163
// easy 3232 find if digit game can be won
// Jan 18, 2025

func canAliceWin(nums []int) bool {
    single := 0
    double := 0
    rest := 0
    ans := false

    for i:=0 ; i<len(nums) ; i++ {
        if nums[i] < 10 {
            single = single + nums[i]
        } else if nums[i] > 9 && nums[i] < 100 {
            double = double + nums[i]
        } else {
            rest = rest + nums[i]
        }
    }

    if single > (double+rest) {
        return true
    } else if double > (single+rest) {
        return true
    }

    return ans
    
}

