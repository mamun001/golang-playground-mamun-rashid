// easy 1437, ELO 1193
// check if all 1's are


func kLengthApart(nums []int, k int) bool {

    ans := true

    first_1_found := false
    last_1_index := -10000
    distance := 10000

    for i:=0;i<len(nums);i++ {
        if nums[i] == 1 {
            if first_1_found == false { // this is the fist 1 of the array
                first_1_found = true
                last_1_index = i
            } else  { // this is NOT the first 1 of the array
                distance = i - last_1_index -1 // as defined in the problem description
                fmt.Println(i,distance)
                if distance < k {
                    return false
                }
                last_1_index = i
            }
        }

    }

    return ans
    
}
