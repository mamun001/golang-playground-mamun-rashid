// easy 1903, ELO 1249
// largest odd number in string

func is_odd(b byte) bool {
    if b == '1' || b == '3' || b == '5' || b == '7' || b =='9' {
        return true
    }
    return false
}

func largestOddNumber(num string) string {


    //fmt.Println(is_odd(num[0]))

    // find the right most odd digit and grab everything to the left of it

    index := -1
    for i:=len(num)-1;i>-1;i-- {
        if is_odd(num[i]) {
            index = i
            break
        }
    }

    if index == -1 {
       return ""
    }

    return num[:index+1]
    
}
