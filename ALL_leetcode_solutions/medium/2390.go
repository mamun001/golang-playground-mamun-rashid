// MEDIUM, 2390, Removing Stars From a String
// from neeetcode's medium youtube playlist , this one has one of the shortest playtime < 5 minutes

// First time , I did this, mybsolution worked , but was too slow, because I was creating and modifying
// strings with [:] at every each iternation, that is too slow! O-squared
// using a slice ("stack)" was much more efficient. That's what I have here.
// I had use to chatgpt to figure why my first alg was too slow.

func removeStars(s string) string {

    ans := ""

    stack := make([]byte,len(s))
    ptr := 0

    for i:=0 ; i<len(s);i++ {
        if s[i] == '*' {
            stack[ptr-1]=0
            ptr--
        } else {
           stack[ptr] = s[i]
           ptr++
        }
    }

    //fmt.Println(stack)
    //fmt.Println(stack[0])

    for i:=0;i<len(stack);i++ {
       if stack[i] != 0 {
          ans = ans + string(stack[i])
       }
    }

    return ans
    
}
