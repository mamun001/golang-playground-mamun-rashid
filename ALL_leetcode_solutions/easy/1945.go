// easy 1945, ELO 1255
// sum of digist of string after convert

func conv_char(b byte) int {

    // a = 96 ("code")
    // a = 96-96=0 => add digits = 0
    // z = 122-96=26 ==> add digits = 8

    if int(b)-96 > 19 {
        return (int(b)-96)%10 +2
    }
    if int(b)-96 > 9 {
        return (int(b)-96)%10 +1
    }
    return int(b)-96
}

func transform(n int) int {
    sum:=0
    q:=n // quotient
    for q > 0 {
      sum=sum+q%10
      q=q/10
    }
    return sum
}

func getLucky(s string, k int) int {

    //fmt.Println(conv_char(s[0]))
    //fmt.Println(conv_char('a'))

    // convert + transform once
    sum:=0
    for i:=0;i<len(s);i++ {
        sum=sum+conv_char(s[i])
    }
    //fmt.Println(sum)
    //fmt.Println(transform(102))

    //transform k-1 times 
    for i:=1;i<k;i++{
        sum=transform(sum)
    }

    return sum
    
}
