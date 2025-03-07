// easy 2485, ELO 1207
// Find the pivot integer

func sum_down(n int) int {
    ans := 0
    for i:=1;i<n+1;i++ {
        ans = ans + i
    }
    return ans
}


func sum_up(n1,n2 int) int {
    ans := 0
    for i:=n1;i<n2+1;i++ {
        ans = ans + i
    }
    return ans
}



func pivotInteger(n int) int {

    if n == 1 {
        return 1
    }

    //fmt.Println(sum_down(3))
    //fmt.Println(sum_up(2,5))

    l:=1
    r:=n
    mid:=0
    

    for l<r {
        mid=(l+r)/2
        //fmt.Println(mid)
        if sum_down(mid) == sum_up(mid,n) {
            return mid
        } else if sum_down(mid) < sum_up(mid,n) {
            l=mid+1
        } else if sum_down(mid) > sum_up(mid,n) {
            r=mid
        }
    }

    return -1
    
}
