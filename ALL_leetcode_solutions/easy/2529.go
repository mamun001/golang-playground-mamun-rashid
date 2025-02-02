// easy 2529, maximum count of positive integer and negative integer
// ELO 1195
// did do binary search because there could be more than 1 zeros
// since maximum number of elements is only 2000, my alg's perf was stil 1ms beats 100% 

func maximumCount(nums []int) int {

    pos := 0
    neg := 0

    for i:=0 ; i<len(nums); i++ {
        if nums[i] < 0 {
            neg++ 
        } else if nums[i] > 0 {
            pos++
        }
    }

    if pos >= neg {
        return pos
    } else if neg > pos {
        return neg
    }
    return 1
    
}

