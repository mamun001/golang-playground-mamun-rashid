
// leetcode easy level
// 448. Find All Numbers Disappeared in an Array


func findDisappearedNumbers(nums []int) []int {
    
    answer := make([]int, 0)
    foo_map := make(map[int]bool)
    
    N := len(nums) 
    fmt.Println(N)
    
    // by default nothing exists; note i here is one more than the index of num
    for i:=1; i <= N ; i++ {
        foo_map[i]=false
    }
    
    // go through the array, whichever is found , turn that key in map to "true"
    for i:=0; i < N; i++ {
        foo_map[nums[i]] = true
    }
    
    // now add the "false" ones into answer slice
    
    for i:=1; i <= N ; i++ {
        if foo_map[i]==false {
            answer=append(answer,i)
        }
    }
    
    return answer
    
}
