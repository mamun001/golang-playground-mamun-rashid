
// leetcode 1047

import (
   "fmt"
)

type stack []byte

func (s stack) Push(v byte) stack {
    return append(s, v)
}

func (s stack) Pop() (stack, byte) {
    // FIXME: What do we do if the stack is empty, though?

    l := len(s)
    return  s[:l-1], s[l-1]
}

func (s stack) PopValueOnly() byte {
    l := len(s)
    if l == 0 {
        return '!'
    } else {
      return s[l-1]
    }
}

func (s stack) PrintStack() {
    l := len(s)
    for i:=0;i<l;i++ {
      fmt.Println(s[i])
    }
}



func removeDuplicates(S string) string {
    //answer := "a"
    
    st := make(stack,0)
    
    // examples using the stack
    //st = st.Push('z')
    //st= st.Push('f')
    //st = st.Push('q')
    //st, p := st.Pop()
    //fmt.Println(p)
    //p := st.PopValueOnly()
    //fmt.Println(p)
    //fmt.Println(string(st)) // PRINT THE WHOLE STACK
    
    
    
    for pos, char := range S {
      //fmt.Printf("character %c starts at byte position %d\n", char, pos)
        
        
        if pos == 0 {
            st = st.Push(byte(char)) // first char, just push to stack
            //fmt.Println("first char ", string(st)) 
        } else {
            if byte(char) == st.PopValueOnly() {
                // duplicate of top char in stack
                st, _ = st.Pop() // pop the top char in stack and don't push this one
                //fmt.Println("duplicate ",string(st)) 
            } else {
                // this char is not a duplicate, push to the stack
                st = st.Push(byte(char))
                //fmt.Println("not duplicate", string(st)) 
            }
          
        } // else    
    } // FOR
    
    //return answer
    return string(st)
}
