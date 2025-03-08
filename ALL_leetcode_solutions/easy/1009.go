// easy 1009, ELO 1235
// Complement of Base 10 integer

// figured out a "math" solution to this bit manipulation problem

func logbase2 (n int) int {
    ans:= 0

    q:= n / 2 

    for q>0 {
        q = q /2
        ans++
    }

    return ans

}

func powerof2 (n int) int {
    ans:=1

    for i:=0;i<n;i++ {
        ans = ans * 2
    }

    return ans
}

func bitwiseComplement(n int) int {

    //fmt.Println(logbase2(n))
    //fmt.Println(powerof2(logbase2(n)+1))

    return powerof2(logbase2(n)+1) - n - 1
    
}
