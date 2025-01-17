// easy 3110 score of a string 
// acceptance rate is 92%



import "fmt"

func abs(i int) int {
    if i < 0 {
        return i * -1
    } else {
        return i
    }
}



func scoreOfString(s string) int {
    sum := 0
    for i := 0 ; i < len(s)-1 ; i++ {
       
       sum = sum + abs(int(s[i]) - int(s[i+1]))
    }

    return sum
}


