// easy 349
// Intersection of Two arrays
// acc rate 75.7
// Binary search

// I HAD LOOK UP part where L = pos +1 and R = pos -1 
// before that I spent 2 hours troubleshooting

// code perf: beats 4 pct cpu, and 7 pct RAM, pretty bad!!
// hopefully it is because of all the print to STDOUS

// we will have to see youtube if I missed something else


import "sort"
import "fmt"


func is_in_slice (slice1 []int, n int) bool {
    ans := false
    for i := 0 ; i < len(slice1) ; i++ {
        if slice1[i] == n {
            return true
        }
    }

    return ans
}

func intersection(nums1 []int, nums2 []int) []int {

    ans := []int{}

    sort.Ints(nums1)
    fmt.Println("nums1 sorted", nums1)
    sort.Ints(nums2)
    fmt.Println("nums2 sorted", nums2)

    LN1 := len(nums1)
    LN2 := len(nums2)


    pos := 0 // inital value, overwritten in first iteration of loop


    // pointers used in Binary Search, initial values
    L := 0
    fmt.Println("L", L)
    R := LN2 -1
    fmt.Println("R", R)


    // items in  num1 , we will look for in nums2
   
    for i_in_nums1 := 0 ; i_in_nums1 < LN1 ; i_in_nums1++ {
        fmt.Println("i_in_nums1" , i_in_nums1)
        fmt.Println("need to find", nums1[i_in_nums1])
        L = 0 // for each new item in nums1 , we start over in nums2
        R := LN2 -1
        for L <= R { // stop when pointers meet
           fmt.Println("")
           fmt.Println("starting inner loop")
           fmt.Println("need to find value", nums1[i_in_nums1])
           pos = (L+R) / 2
           fmt.Println("pos in nums2", pos)
           fmt.Println("value in nums2", nums2[pos])
           if nums1[i_in_nums1] == nums2[pos] {
               fmt.Println("found match", nums1[i_in_nums1])
               fmt.Println("breaking inner loop")
               if is_in_slice (ans,nums1[i_in_nums1] ) != true {
                  ans = append (ans, nums1[i_in_nums1]) // found a match
               }
               break // break inner loop and move to next item in num1 and look for that
           } else if nums1[i_in_nums1] < nums2[pos] { // look left
               fmt.Println("value to find is less than what is in current pos")
               fmt.Println("L R before change", L,R)
               R = pos -1 
               fmt.Println("moved to Left in nums2, L R", L,R)
               fmt.Println("need to find value", nums1[i_in_nums1])
           } else if nums1[i_in_nums1] > nums2[pos] { // look to the right
               fmt.Println("value to find is more than what is in current pos")
               fmt.Println("L R before change", L,R)
               L = pos + 1
               fmt.Println("moved to Right in nums2, L R", L,R)
               fmt.Println("need to find value", nums1[i_in_nums1])
           } 
        } // for inner
        fmt.Println("ended inner loop: i_in_nums1 ", i_in_nums1 )
        fmt.Println("")
    } // for outer
    return ans
  } //func

