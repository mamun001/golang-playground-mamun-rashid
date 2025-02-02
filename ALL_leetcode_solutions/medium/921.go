// 921 MEDIUM Minimum Add to Make Parentheses Valid
// Clever algorithm by me!
// CPU > 100% and RAM > 64%

// Jan 28, 2025


func minAddToMakeValid(s string) int {


   L := 0 // (
   R := 0 // )

   // R closes existing L, but
   // L does not close existing R
   // return sum of all the leftovers
   // although , i think, only L or R can be non-zero in the end
   
    for i:=0;i<len(s);i++ {
        if s[i] == '(' {
            L++
        } else if s[i] == ')' && L > 0 {
            L--
        } else if s[i] == ')' && L == 0 {
            R++
        }
    }



    return L+R
    
}
