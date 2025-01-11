// easy 1021
// 84.6 pct acc rate
// Remove Outermost Parentheses

// we are supposed to use stack
// but found a clever algorithm where keep "balace" and when balance hits 0, i know it is the end
// of a "valid parenthesis string"

// that was ok, because my code beat 41% in cpu and 47% in RAM

// BUT, that made the 2nd of the code difficult, spend a lot of time troubleshooting

import "fmt"

func removeOuterParentheses(s string) string {

    ans := ""
    BAL := 0 // "balnace of left minus right"
    indexes_to_remove := []int{0} // first is always removed

    //fmt.Println(len(s))

    for i:=0 ; i < len(s) ; i++ {
        if s[i] == '(' {
            BAL++
            //fmt.Println(BAL)
        } else if s[i] == ')' {
            BAL--
            //fmt.Println(BAL)
        }
        if i !=0 && i != len(s)-1 && BAL == 0 { // WE ARE AT THE END OF VALID PARENTHESIS STRING
            indexes_to_remove = append(indexes_to_remove,i)
            indexes_to_remove = append(indexes_to_remove,i+1)
        }
    }
    indexes_to_remove = append(indexes_to_remove,len(s)-1) // last one is always removed
    //fmt.Println("indexes_to_remove:",indexes_to_remove)

    // testing how to remove nth char from a string
    //s2 := "abcdefghijklmnopqrst"
    //fmt.Println("s2: ",s2)
    //fmt.Println(s2[0:2]) // char 0 to 1 inclusive
    //fmt.Println(s2[3:8]) // char 3 to 7 inclusive
    //s3 := s2[0:2] + s2[3:8] // remove char 2
    //fmt.Println(s3)
    //n := 2 // remove char 2
    //s3 := s2[0:n]+s2[n+1:len(s2)]
    //fmt.Println(s3)

    // can I copy a string?
    //s2 := s
    //fmt.Println(s2)

    // if indexes_to_remove has 0,5
    //s3 := s2[1:5] // take out 0th and 5th, keep 1st-4th
    //fmt.Println("s3: ",s3)

    
    
    for i:=0 ; i <=len(indexes_to_remove)-2 ; i=i+2{
        //fmt.Println(i)
        //fmt.Println(indexes_to_remove[i],indexes_to_remove[i+1])
        //ans = ans + s2[test[i]+1:test[i+1]]
        ans = ans + s[indexes_to_remove[i]+1:indexes_to_remove[i+1]] // TAKE OUT ONE VALID PARENTESIS STRING and append
    }
    //fmt.Println(s2)

    return ans
    
}
