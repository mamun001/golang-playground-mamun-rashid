
import "strconv"


func countDistinctIntegers(nums []int) int {

    myint := 12345678
    fmt.Println(myint) // 12345678
    s1 := strconv.Itoa(myint)
    fmt.Println(s1) // 12345678

    // reverse the string so we can reverse the int effectively
    s_rev := ""
    for i:=len(s1)-1; i > -1 ; i-- {
        s_rev = s_rev + string(s1[i])
    }
    fmt.Println(s_rev) // 87654321
    reversed_int,_ := strconv.Atoi(s_rev)
    fmt.Println(reversed_int) // 87654321


    

    return 0
    
}
