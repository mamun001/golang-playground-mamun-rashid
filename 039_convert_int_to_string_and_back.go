import "strconv"


func countDistinctIntegers(nums []int) int {

    myint := 12345678
    fmt.Println(myint) // 12345678
    s1 := strconv.Itoa(myint)
    fmt.Println(s1) // 12345678
    anotherint,_ := strconv.Atoi(s1) // 12345678
    fmt.Println(anotherint)
    

    return 0
    
}

