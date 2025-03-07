// using this space (TWO SUM) to practice merge sort


func merge_sort(ar []int) []int {
    if len(ar) <= 1 {
        return ar
    }

    mid := len(ar) / 2
    left := ar[:mid]
    right := ar[mid:]

    left_sorted := merge_sort(left)
    right_sorted := merge_sort(right)

    return merge(left_sorted,right_sorted)


}

func merge(a,b []int) []int {

    ans := []int{}
    pa := 0
    pb := 0

    //fmt.Println(a[pa])
    //fmt.Println(b[pb])
    

  for pa<len(a) && pb<len(b) {
    if a[pa] < b[pb] {
        ans = append(ans,a[pa])
        pa++
    } else {
        ans = append(ans,b[pb])
        pb++
    } 
  }

  // add remaining ones
  ans = append(ans,a[pa:]...)
  ans = append(ans,b[pb:]...)

  return ans

}


func twoSum(nums []int, target int) []int {
    s := make([]int, 2)

    a := []int{1,9,4,5,2,8,4,7,2}
    //a := []int{9,1}

    //b := []int{1,3,5}
    //c := []int{2,4,6,8}
    //fmt.Println(merge(b,c)) // testing the merge func

    ans := merge_sort (a)
    fmt.Println(ans)

    return s
}
