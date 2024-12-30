// easy 3280 
// 88 pct acc rate

// gave uo and went to youtube for a more elegant solution
// but there were none in Golang
// even chatgpt could not help

// Python and Java has libabries that help, golang does not/
// That explains why the acc rate is so high

// my alg is is as good as any


import "fmt"
import "strconv"



func convert_to_bin_string(s string) string {
    // edge cases
    if s == "0" {
        return "0"
    }
    

    ans := ""
    num, _ := strconv.Atoi(s)
    mod1 := 0
    quo1 := num
    //fmt.Println(a)
    for i := 0 ; quo1 > 0 ; i++ {
        mod1 = quo1 % 2
        quo1 = quo1 / 2
        ans = strconv.Itoa(mod1) + ans
        //fmt.Println(num,"q ",quo1,"mod ",mod1,"ans: ", ans)
        
        
    }

    //return reverse_string(ans)
    return ans
}


func convertDateToBinary(date string) string {
    year := date[0:4]
    month := date[5:7]
    dom := date[8:10]
    //fmt.Println(year,month,dom)
    //fmt.Println(convert_to_bin_string(year))
    //fmt.Println(convert_to_bin_string(month))
    //fmt.Println(convert_to_bin_string(dom))
    return convert_to_bin_string(year)+"-"+convert_to_bin_string(month)+"-"+convert_to_bin_string(dom)
    
}



