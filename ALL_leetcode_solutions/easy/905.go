// easy 905, ELO 1178
// sort array by parity

// JAN 22, 2025
// Used two pointers


func sortArrayByParity(nums []int) []int {

    z:=len(nums)-1

    for i:=0 ; i <= z ; i++ {
        //fmt.Println(i)
        if nums[i] % 2 == 1 { // odd
          for z > i {
            //fmt.Println(z)
            if nums[z] % 2 == 0 {
                //fmt.Println("found even, z, numz", z, nums[z])
                nums[i],nums[z] = nums[z],nums[i]
                z--
                break
            }
            z--
          } // for
        } // if 
    } // for i

    return nums
    
}
