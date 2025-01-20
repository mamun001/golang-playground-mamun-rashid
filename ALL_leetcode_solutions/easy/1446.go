// easy 1446 consecutive characters
// ELO 1165 or so
// JAN 18, 2025

// cataching all the edge cases is hard here, feel like code fails on a new one when code is changed
// that explains 60 pct acc rate
// DOES NOT explain the low ELO

import "fmt"

func maxPower(s string) int {

    repeat_count := 1
    power := 1 // most repeat count thus far

    for i,v := range s {

        
        if i == 0 {
            continue // first one does not matter
        }
        
        if v == rune(s[i-1])  { // repeat starts or continues
              fmt.Println("match")
              repeat_count = repeat_count + 1
              //fmt.Println("repeat",repeat_count)
              if repeat_count > power {
                  power = repeat_count // record broken, new max
              } else { // repeat continues but max/power record was not broken
                  power = power + 0
               } // else
        } else { // no longer repeats, new char showed up
            repeat_count = 1
        }
        // end of each i
    } // for

    return power
    
} // func
