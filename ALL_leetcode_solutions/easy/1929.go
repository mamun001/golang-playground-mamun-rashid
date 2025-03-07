// 1929 easy concatenation of array

// THE WHOLE PROBLEM IS JUST KNOWING YOU CAN RETURN A SLICE!!!
// Otherwise, this is impossible to do in Golang
// Confirmed by chatgpt

// For this one, I had to look up the solution

// DEC 22, 2024


func getConcatenation(nums []int) []int {
    // DEC 22, 2024

    nums2 := nums
    result := append (nums,nums2...)
   
    return result
}
