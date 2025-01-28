// MEDIUM 2414 length of the longest alphabetical ...
// ELO 1221
// Spent 4 Hours debugging edge cases and bugs in code!

// My alg: CPU > 25% , RAM > 80%, so not bad!

// Jan 27, 2025

// My alg is pretty, but it works
// Its convoluted because of so many tough edge cases!
// I should look up the answer to see what I missed


func longestContinuousSubstring(s string) int {

    max :=0
    this_streak :=0
    broken_once := false
    last_broken_at :=0

    if len(s) == 1 {
        return 1
    }

    for i:=0;i<len(s);i++ {
        //fmt.Println("i",i,"max",max,"rune",rune(s[i]))
        if i == 0 {
            max = 1
            this_streak = 1
        } else if i > 0 && (s[i] == s[i-1]+1) { // streak on
            this_streak++
        } else if i > 0 && (s[i] != s[i-1]+1) { // streak broken
               broken_once = true
               last_broken_at = i
               //fmt.Println("broken",i,s[i],s[i-1])
               if this_streak > max {
                max = this_streak
               }
               this_streak = 1 // new start
        }
    } // for

    if broken_once == false {
        //fmt.Println("not broken")
        return len(s)
    }

    if last_broken_at != len(s)-1 {
        if max > len(s) - last_broken_at {
            return max
        } else {
        return len(s) - last_broken_at
        }
    }
            
    return max
    
} // func

