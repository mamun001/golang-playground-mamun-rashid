//
// Input: command = "G()(al)"
// Output: "Goal"
// Explanation: The Goal Parser interprets the command as follows:
// G -> G
// () -> o
// (al) -> al
// The final concatenated result is "Goal".


func interpret(command string) string {
    answer := "abc"
    
    input_byte_array :=[]byte(command)
    vals := make([]byte, 0)
    
    open_p := false
    open_l := false
    
    for i:=0;i<len(input_byte_array);i++ {
        if input_byte_array[i] == 'G' {
            vals = append(vals, 'G')
        } else if input_byte_array[i] == '(' {
            open_p = true
        } else if input_byte_array[i] == 'l' {
            open_l = true
            open_p = false
        } else if input_byte_array[i] == ')' {
            if open_p == true {
                vals = append(vals, 'o')
                open_p = false
            } else if open_l == true {
                vals = append(vals, 'a')
                vals = append(vals, 'l')
                open_l = false
            }
            
        }
             
    }
    
    answer = string(vals)
    
    
    
    return answer
    
}
