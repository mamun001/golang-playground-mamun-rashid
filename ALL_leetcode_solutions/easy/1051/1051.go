// 80 percent acc rate

// 1051 Height Checker 

// DEC 25, 2024


//import "fmt"
//import "sort" // we don't even need this in leetcode to use sort package

func heightChecker(heights []int) int {

   ans := 0

   duplicate_slice := make([]int, len(heights))
   copy(duplicate_slice,heights)
   sort.Ints(heights)

   //fmt.Println(heights)
   //fmt.Println(duplicate_slice)

   for i,_ := range heights {
    if heights[i] != duplicate_slice[i] {
        ans++
    }
   }

   return ans
 
}



