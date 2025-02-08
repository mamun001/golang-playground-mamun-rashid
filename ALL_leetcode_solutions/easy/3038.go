// easy 3038, ELO 1202
// maximum number of operayions with the same score I

func maxOperations(nums []int) int {

    sum := nums[0]+nums[1]
    ans := 1
    streak_broken := false
    m := len(nums)
    loop := 0
    this_sum := 0
    next := 0
    next_next := 0

    if m % 2 == 0 {
        loop = (m-2)/2
    } else {
        loop = (m-3)/2
    }
    fmt.Println("loop",loop)

    for i:=1;i<=loop && streak_broken == false;i=i+1{
        next = 1+2*(i-1)+1
        next_next = 1+2*(i-1)+2
        fmt.Println("indexes",next,next_next)
        this_sum = nums[next] + nums[next_next]
        if this_sum == sum {
            ans++
        } else {
            streak_broken = true
        }
        
    }




    return ans
    
}

