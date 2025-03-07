// MEDIUM 912, sort and array, no ELO
// acc rate 57%
// part of Neetcode 250

// my solution took too long twice!
// my alg AND merge sort! both took too long
// SO, I asked chatgpt
// chatgpt solution worked 
// I looked at the difference
// in lines 23-30, you will see that I calling merge_sort twice!!!
// that was the problem
// once I fixed that, even mine worked


// MINE

func merge_sort (ar []int) []int {
    if len(ar) <= 1 {
        return ar
    }

    mid := len(ar)/2

    // I AM SORTING TWICE , that's why I solution took too long!
    //left := merge_sort(ar[:mid])
    //right := merge_sort(ar[mid:])
    //return merge(merge_sort(left),merge_sort(right))

    left := ar[:mid]
    right := ar[mid:]
    return merge(merge_sort(left),merge_sort(right))

}


/*
// FROM CHATGPT
func sort2(nums []int) []int {
    if len(nums) <= 1 {
        return nums
    }

    mid := len(nums) / 2
    left := sort2(nums[:mid])
    right := sort2(nums[mid:])

    return merge(left, right)
}
*/


// MINE

func merge(a , b []int) []int {

   ans := []int{}

   pa := 0
   pb := 0

   for pa < len(a) && pb <len(b) {
     if a[pa] < b[pb] {
        ans = append(ans,a[pa])
        pa++
     } else {
        ans = append(ans,b[pb])
        pb++
     }
   }

   // left overs
   ans = append(ans,a[pa:]...)
   ans = append(ans,b[pb:]...)

   return ans
}



/*
// FROM CHATGPT

func merge(left, right []int) []int {
    result := make([]int, 0, len(left)+len(right))
    i, j := 0, 0

    for i < len(left) && j < len(right) {
        if left[i] < right[j] {
            result = append(result, left[i])
            i++
        } else {
            result = append(result, right[j])
            j++
        }
    }

    result = append(result, left[i:]...)
    result = append(result, right[j:]...)

    return result
}
*/


func sortArray(nums []int) []int {

    sorted := merge_sort(nums)

    return sorted
    
}
