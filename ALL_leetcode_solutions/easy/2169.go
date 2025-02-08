// easy 2169 , ELO 1199 
// count operations to obtain zero

func countOperations(num1 int, num2 int) int {

    ans:= 0

    for num1 != 0 && num2 != 0 {
        ans++
        if num1 >= num2 {
            num1 = num1 -num2

        } else {
            num2 = num2 - num1
        }
    }

    return ans
    
}
