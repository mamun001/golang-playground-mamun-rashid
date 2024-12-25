package main

import "fmt"
import "sort"


// sort slice using "sort" package already available.
// in leetcode you don't even have to import!
// but in my mac, I have "import"

func main() {
      s := []int{1,20,2,34,50,2,500,1000,3,456}  // a slice of integers
      sort.Ints(s)	 // sorts in place. s itself gets newly sorted items
      fmt.Println(s)

}
