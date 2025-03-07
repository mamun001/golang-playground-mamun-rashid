// easy 1952, ELO 1203
// Three divisors

func isThree(n int) bool {

    if n <=3 {
        return false
    }

    //ans := false

    divisors := []int{1,n}
    //fmt.Println(divisors)

    for i:=2;i<(n/2)+1;i++ {
        if n % i == 0 {
            divisors = append(divisors,i)
        }
    }

    return len(divisors) == 3
    
}
