// easy, number 27
// acc rate 58.9%

// DEC 29, 2024

// GOT HELP FROM YOUTUBE , because my algorithms (I trued two) were too complex and did not work
// I spend at least 2 hours before going to Youtube

// even after that, I could not code it right, even some of my own test cases would fail

// so, I had convert youtube's Python code to Golang
// https://www.youtube.com/watch?v=Pcd1ii9P9ZI


// I MUST TRY THIS AGAIN ON MY WON in A MONTH OR TWO

// IN TOTAL, I spend 3 hours on this problem!

import "fmt"

func removeElement(nums []int, val int) int {

    k := 0

    for i := 0 ; i < len(nums); i++ {
        if nums[i] != val {
            nums[k] = nums[i]
            k = k+1
        }
    }
    
    return k

