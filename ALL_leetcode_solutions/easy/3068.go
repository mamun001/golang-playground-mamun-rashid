// easy 3068, ELO 1203
// Distribute Elements Into Two Arrays

func resultArray(nums []int) []int {

    ans := []int{}

    arr1 := []int{nums[0]}
    arr2 := []int{nums[1]}
    //fmt.Println(nums[1])

    for i:=2;i<len(nums);i++ {
      if arr1[len(arr1)-1] > arr2[len(arr2)-1] {
          arr1 = append(arr1,nums[i])
      } else {
          arr2 = append(arr2,nums[i])
      }
    }


   //fmt.Println(arr1)
   //fmt.Println(arr2)

   ans = append(arr1,arr2...)

    return ans
    
}
