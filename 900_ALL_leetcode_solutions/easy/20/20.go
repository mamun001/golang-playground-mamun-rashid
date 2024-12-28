// 4th easiest problem according to chatgpt

// PROBLEM #20 , easy level

// 41.5 pct acc rate

// DEC 28 , 2024

// test cases did not show this, but looks like you can multiple same brackets open
// for example (([])) is valid input

//import "fmt"

func isValid(s string) bool {

    ans := true // all good , by default

    first_brac_open := 0
    second_brac_open := 0
    third_brac_open := 0
    first_brac_close := 0
    second_brac_close := 0
    third_brac_close := 0


    stack_of_opens := []int{} // we keep track of order of opened bracked using a stack, 1 is first , 2 is second, 3 is third bracket

    // first loop to do simple checks of unclosed brackets

    for _,v := range s {

        if byte(v) == '('  {
           //fmt.Println("first bracket open")
           first_brac_open++
        }
        if byte(v) == ')'  {
           //fmt.Println("first bracket closed")
           first_brac_close++
        }

        if byte(v) == '{'  {
           //fmt.Println("second bracket open")
            second_brac_open++
        }
        if byte(v) == '}'  {
           //fmt.Println("second bracket closed")
           second_brac_close++
        }

        if byte(v) == '['  {
           //fmt.Println("third bracket open")
            third_brac_open++
        }
        if byte(v) == ']'  {
           //fmt.Println("second bracket closed")
           third_brac_close++
        }

    } // for loop
    

    if first_brac_open != first_brac_close || second_brac_open != second_brac_close || third_brac_open != third_brac_close {
        return false
    }
    
    // // second loop to check for wrong brackets closed

    fmt.Println("second loop")
    for _,v := range s {
        switch byte(v)   {
          case '(': // we have to PUSH to the stack
             //fmt.Println("pushing 1 to stack")
             stack_of_opens = append(stack_of_opens,1)
          case ')':
            if len(stack_of_opens) == 0 || stack_of_opens[len(stack_of_opens)-1] != 1 { // wrong bracket is getting closed 
                 //fmt.Println("wrong bracket is getting closed",")")
                 return false
            } else { // correct bracket is getting closed, we have to POP the last one
                //fmt.Println("correct bracket is getting close, popping",")")
                stack_of_opens = stack_of_opens[:len(stack_of_opens)-1] // does it work?? yes
            } // else

          case '{': // we have to PUSH to the stack
             //fmt.Println("pushing 2 to stack")
             stack_of_opens = append(stack_of_opens,2)
          case '}':
            if len(stack_of_opens) == 0 || stack_of_opens[len(stack_of_opens)-1] != 2 { // wrong bracket is getting closed 
                 //fmt.Println("wrong bracket is getting closed","}")
                 return false
            } else { // correct bracket is getting closed, we have to POP the last one
                //fmt.Println("correct bracket is getting close, popping","}")
                stack_of_opens = stack_of_opens[:len(stack_of_opens)-1] // does it work?? yes
            } // else

          case '[': // we have to PUSH to the stack
             //fmt.Println("pushing 3 to stack")
             stack_of_opens = append(stack_of_opens,3)
          case ']':
            if len(stack_of_opens) == 0 || stack_of_opens[len(stack_of_opens)-1] != 3 { // wrong bracket is getting closed 
                 //fmt.Println("wrong bracket is getting closed","]")
                 return false
            } else { // correct bracket is getting closed, we have to POP the last one
                //fmt.Println("correct bracket is getting close, popping","]")
                stack_of_opens = stack_of_opens[:len(stack_of_opens)-1] // does it work?? yes
            } // else


        } // switch 
    } // for

    return ans
}

