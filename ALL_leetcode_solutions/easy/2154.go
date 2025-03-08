// easy 2154, ELO 1236
// keep multiplying found values by two

func max(ar []int) int {
    
    max := 0
    for i:=0;i<len(ar);i++ {
        if ar[i] > max {
            max = ar[i]
        }
    }
    return max
}


func is_in(ar []int, n int) bool {
    
    for i:=0;i<len(ar);i++ {
        if ar[i] == n {
            return true
        }
    }
    return false
}

func findFinalValue(nums []int, original int) int {

    n := original
    m := max(nums)

    for n <= m {
        if is_in(nums,n) {
            n = n*2
        } else {
            return n
        }
    }
    
    return n
    
}
