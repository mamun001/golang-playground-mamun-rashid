// MEDIUM 128, no ELO, part of NC 250
// acc rate 47%
// cheated though, used sort package which is Ologn complexity
// but since Go sort package is so fast that it works

import "sort"

func longestConsecutive(nums []int) int {


    if len(nums) == 0 {
        return 0
    }

    if len(nums) == 1 {
        return 1
    }

    sort.Ints(nums)
    //fmt.Println(nums)
    ans := 0
    current_seq := 1
    

    for i:=1 ; i<len(nums);i++ {
        if nums[i] == nums[i-1] {
            continue
        }
        if nums[i] == nums[i-1] + 1 {
            current_seq++
            //fmt.Println(nums[i],current_seq)
        } else {
            if current_seq > ans {
                ans = current_seq
            }
            current_seq = 1
        }
    }

    if current_seq > 0 {
        if current_seq > ans {
                ans = current_seq
            }
    }

       return ans
    
}
