
// #283 Leetcode easy level

func run_one_loop(nums []int) {
    
  L := len(nums)
    
  for i:=0; i < L; i++ {
        if nums[i] == 0 {
            // bring every item LEFT 1 slot
            for j:=i; j < L-1 ; j++ {
                nums[j] = nums[j+1]
            }
            nums[L-1] = 0
        }
    }
}



func moveZeroes(nums []int)  {
    
    L := len(nums)
    
    leading_zeros := 0
    
    for i:=0; i<L ; i++ {
        if nums[i] != 0 {
            // no more leading zeros
            leading_zeros = i        
        }
    }
    fmt.Println (leading_zeros)
    
    // we have to run the loop n-1 times where n is the number os leading zero's
    
    for i:=0; i < leading_zeros ; i++ {
        run_one_loop(nums)
    }
}
