// easy 2032 ELO 1270
// Two out of three

func is_in(ar []int, n int) bool {
    ans := false
    for i:=0;i<len(ar);i++ {
        if ar[i] == n {
            return true
        }
    }
    return ans
}

func rem_dup(ar []int) []int{
    ans :=[]int{}
    for i:=0;i<len(ar);i++ {
        if is_in(ans,ar[i]) == false {
            ans = append(ans,ar[i])
        }
    }
    return ans
}

