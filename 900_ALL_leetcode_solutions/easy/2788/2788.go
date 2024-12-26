// 74%

// easy level

// 2788 Split Strings by Separator

// DEC 25, 2024

// TOUGH PROBLEM!, SPENT 90+ on it

// my algorithm must not be very good, because my time beat only 8% and memory only 13%


//import "fmt"

func splitWordsBySeparator(words []string, separator byte) []string {

    ans := []string{}
    making_a_new_string := ""

    for i,_ := range words { // one word at a time
        //fmt.Println(v)
        making_a_new_string = "" // start of a new element in "words" is also considered a boundary just like the separator
        for j,x := range words[i] { // one character at time
           //fmt.Println(rune(x))
           //fmt.Println(rune(separator))
           

           if rune(x) != rune(separator) { // still no sight of the separator
            //fmt.Println("match")
            making_a_new_string = making_a_new_string + string(rune(x))
           } else {
            // separator has come
            if making_a_new_string != "" {
               ans = append(ans,making_a_new_string)
               making_a_new_string = ""
            }
            //fmt.Println(making_a_new_string)
           } // else

           if j == len(words[i]) -1 { // we just ended an element in words array
              if making_a_new_string != "" {
                 ans = append(ans,making_a_new_string)
                 making_a_new_string = ""
              }
           }


        } // inner for loop
    } // outer for loop

    return ans
} // func


