// 209, MEDIUM, Minimum Size Subarray Sum
// my alg worked 85% of the time
// then I had to go to you tube


// my submissions kept failing further and furter into testing cases
// after 2+ hours of debguuing, went to youtube
// the key thing I missed that , if sum is still a good sum, we move the Left pointer
// that's the part I missed from my alg



// finally DONE

func minSubArrayLen(target int, nums []int) int {

    N := len(nums)
    cur_sum:= nums[0]
    min_len := 100001
    l:=0
    r:=0


    
    for l<N && r<N {
        //fmt.Println(nums[l],nums[r])
        if cur_sum >= target && r-l+1 < min_len {
            min_len = r-l+1
            cur_sum = cur_sum - nums[l]  
            l++        
        } else if cur_sum >= target && r-l+1 >= min_len {
            cur_sum = cur_sum - nums[l]
            l++   
        } else if cur_sum < target && r != N-1 {
            r++
            cur_sum = cur_sum + nums[r] 
        } else if cur_sum < target && r == N-1 {
            cur_sum = cur_sum - nums[l] 
            l++ 
        }
    }


    if min_len == 100001 {
        return 0
    }


    return min_len
    
}
