// easy 2864, ELO 1238
// maximum odd binary binary number

func maximumOddBinaryNumber(s string) string {

    ans:= "1"

    ones:=0
    zeros:=0

    for i:=0;i<len(s);i++ {
        if s[i] == '1' {
            ones++ 
        } else {
            zeros++
        }
    }


    for i:=0;i<zeros;i++ {
        ans = "0" + ans
    }

    for i:=0;i<ones-1;i++ {
        ans = "1" + ans
    }



    return ans
    
}
