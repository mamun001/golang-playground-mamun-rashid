// easy 922, ELO 1173
// sort array by parity

import "sort"

func sortArrayByParityII(nums []int) []int {
    
    sort.Ints(nums)
    for i:=0 ; i < len(nums)-1; i++ {
       if i % 2 == 0 {
        if nums[i] % 2 == 1 {
            for j:=i+1;j<len(nums);j++{
                if nums[j] % 2 == 0 {
                    nums[i],nums[j] = nums[j],nums[i]
                    break
                }
            }
        }
       }
       if i % 2 == 1 {
        if nums[i] % 2 == 0 {
            for j:=i+1;j<len(nums);j++{
                if nums[j] % 2 == 1 {
                    nums[i],nums[j] = nums[j],nums[i]
                    break
                }
            }
        }
       }


    }

    //fmt.Println(nums)


    return nums
    
}

