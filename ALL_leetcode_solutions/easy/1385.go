// easy 1385, ELO 1234
// although problem description is horrible, coded it up in 15 minutes
// ran the first time
// did nit use Binary Search
// alg perf: 100% + 74% = pretty darn good


func abs (n int) int {
    ans := n
    if n < 0 {
        ans = n * -1
    }

    return ans
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {

    ans := 0
    m := len(arr1)
    n := len(arr2)
    far := true

    for i:=0;i<m;i++ {
        far = true
        for j:=0;j<n && far == true;j++ {
            if abs(arr1[i]-arr2[j]) <= d {
                far = false
            }
        }
        if far == true {
            ans++
        }
    }

    return ans
    
}
