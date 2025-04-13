// leetcode problem number 3, MEDIUM
// no ELO, acc rate 37%
// no help needed, Apil 12th
// with debugging , the whole thing took < 15 minutes!

func lengthOfLongestSubstring(s string) int {


    // edge cases that are simple
    if len(s) < 2 {
        return len(s)
    }

    longest := 1

    l:=0
    r:=1

    map_letters := make(map[byte]int)
    map_letters[s[0]]++

    for r < len(s) {
        if map_letters[s[r]] == 0 {
            // dup still not found
            map_letters[s[r]]++
            r++ // only move r pointer
            continue
        } else {
            // dup found
            if r-l > longest {
                longest = r-l
            }
            map_letters[s[l]] = 0 // now that we are about to l 1 to the right, we must that letter from the map
            l++ // keep r pointer where it is
        }

    }


   // when the last char in the string is still not a repeat
   if r-l > longest {
                longest = r-l
   }



    return longest
    
}
