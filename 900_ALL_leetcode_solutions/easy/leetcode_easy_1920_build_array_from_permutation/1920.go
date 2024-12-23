

// easy 1920 Build Array from Permutation
// DEC 22, 2024


func buildArray(nums []int) []int {
    output := []int{} // this will eventually be returned with the right values

    // fill it up with semi-random values os that "output" slice has the right number of items
    for _,v := range nums {
        output = append(output,v)
    }
    
    // This formual you can get directly from the question
    for i,_ := range nums {
        output[i] = nums[nums[i]]
     }


    return output
}

