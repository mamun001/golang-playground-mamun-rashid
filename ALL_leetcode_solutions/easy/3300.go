// easy 3300, minimum element after 
// ELO 1181
// JAN 23, 2025

func sum_of_digits (n int) int {
    ans := 0
    digits := []int{}
    q := n
    for q != 0 {
        digits = append(digits, q % 10 )
        q = q / 10
    }
    
     for i:=0; i < len(digits); i++ {
        ans = ans + digits[i]
     }

     return ans
}

func min(arr []int) int {
    min := 10001
    for i:=0; i < len(arr); i++ {
        if arr[i] < min {
            min = arr[i]
        }
    }
    return min
}


func minElement(nums []int) int {

    //ans := 0

    //fmt.Println(sum_of_digits(6666))
    for i:=0; i<len(nums) ; i++ {
        nums[i] = sum_of_digits(nums[i])
    }

    fmt.Println(nums)
    return min(nums)
    
}

