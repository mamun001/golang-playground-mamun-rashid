
func runningSum(nums []int) []int {
    
    //var result [len(nums)]int
    s := make([]int, len(nums))
    var running_sum int = 0
    
    for i := 0; i < len(nums); i++ {
        for j := 0; j <= i ; j++ {
            running_sum = running_sum + nums[j]
        }
        s[i] = running_sum 
        running_sum = 0       
    }
    
    return s
    
}
