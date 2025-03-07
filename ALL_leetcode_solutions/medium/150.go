// MEDIUM 150, part of neetcode 250
// Evaluate Reverse Polish Notation

import "strconv"

func evalRPN(tokens []string) int {

    
    n := 0
    a := 0
    b := 0


    stack:=[]int{}

    for i:=0;i<len(tokens);i++ {
        //fmt.Println(tokens[i])
        //fmt.Println(rune(tokens[i]))

        if tokens[i] == "+" {
            b = stack[len(stack)-1]
            a = stack[len(stack)-2]
            stack = stack[0:len(stack)-2]
            stack = append(stack,a+b)
        } else if tokens[i] == "-" {
            b = stack[len(stack)-1]
            a = stack[len(stack)-2]
            stack = stack[0:len(stack)-2]
            stack = append(stack,a-b)
        } else if tokens[i] == "*" {
            b = stack[len(stack)-1]
            a = stack[len(stack)-2]
            stack = stack[0:len(stack)-2]
            stack = append(stack,a*b)
        } else if tokens[i] == "/" {
            b = stack[len(stack)-1]
            a = stack[len(stack)-2]
            stack = stack[0:len(stack)-2]
            stack = append(stack,a/b)
        } else {
            n,_ = strconv.Atoi(tokens[i])
            stack = append(stack,n)
        }

    }

    return stack[0]
    
}
