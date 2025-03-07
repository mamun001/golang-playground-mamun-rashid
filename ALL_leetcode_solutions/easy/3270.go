// easy 3270, ELO 1205
// Find the key of the numbers

func min (a,b,c int) int {
    if a <=b && a <= c {
        return a
    }
    if b <=c && b <= a {
        return b
    } else {
        return c
    }
}

func generateKey(num1 int, num2 int, num3 int) int {

    //fmt.Println(min(10,1,2))

    d1 := min(num1/1000,num2/1000,num3/1000)
    fmt.Println(d1)
    d2 := min((num1%1000)/100,(num2%1000)/100,(num3%1000)/100)
    fmt.Println(d2)
    d3 := min((num1%100)/10,(num2%100)/10,(num3%100)/10)
    fmt.Println(d3)
    d4 := min((num1%10)/1,(num2%10)/1,(num3%10)/1)
    fmt.Println(d4)

    ans := d1*1000+d2*100+d3*10+d4

    return ans
    
}
