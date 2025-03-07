// easy 2148 , ELO 1202
// count elements with strictly smaller
// brute force works



func countElements(nums []int) int {

    //sort.Ints(nums)

    ans:=0
    

    //winners :=[]int{}

    for i:=0;i<len(nums);i++ {
        smaller := false
        larger := false
        for j:=0;j<len(nums);j++ {
          if i != j && nums[i] > nums[j] {
            smaller = true
          }
          if i != j && nums[i] < nums[j] {
            larger = true
          }
        }
        if smaller && larger {
            ans++
        }

    }

    return ans
    
}
