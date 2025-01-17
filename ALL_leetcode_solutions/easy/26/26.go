// easy , #26
// 6th easiest per chat gpt
// 58.7 acc rate

// DEC 28, 2024

// Had to get help from youtube only to see how you modify incoming slice
// Otherwise , I had the algorithm right

//import "fmt"

func removeDuplicates(nums []int) int {

    // nums[0] = 1 testing to see if this works
    // ums[1] = 2

    left_pointer := 1 // 0th element will never change, THIS IS WHERE WE ADD next non-duplicate
    right_pointer :=1 // we start in the same place , THIS IS WHERE WE "look"

    for i := 1 ; i < len(nums); i++ {
        if nums[i] == nums[i-1] { // duplicate
           right_pointer++ // since a duplicate, we keep lookimg
           //fmt.Println(i, nums[i], "duplicate")
        } else { // NEW NUMBER
            //fmt.Println(i, nums[i], "not a duplicate")
            nums[left_pointer] = nums[i] // we the new number in first available slot on the left
            //fmt.Println(nums[0],nums[1],nums[2])
            right_pointer++
            left_pointer++

        }
    }

    k := left_pointer

    return k
    
}

