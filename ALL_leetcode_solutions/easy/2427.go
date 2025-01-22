// easy 2427, ELO 1172
// Number of Common Factors

// Jan 20, 2025

func commonFactors(a int, b int) int {

    ans := 0

    for i:=1 ; i<=b && i <=a ; i++ {
        if a % i == 0 && b % i == 0 {
            ans++
            //fmt.Println(i)
        }
    }

    return ans
    
}
