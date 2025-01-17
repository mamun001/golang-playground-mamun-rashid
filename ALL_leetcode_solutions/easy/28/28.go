// easy leetcode number: 28

// acc rate 44.1%

// DEC 29, 2024

// algorithm is kind of messy, but it takes less than a ms and beats 24% of the other codes in RAM usage

// took me 90 minutes to write and debug, i bet youtube has better algorithm

import "fmt"

func strStr(haystack string, needle string) int {

ans := -1 // default answer is that needle does not exist in the haystack
start_of_matching := 0 // as we go through the haystack, we keep track of our "start" index of current matching with the needle
broken := false // used in line 36, when we are done with one search of needle, we need to know if search succeeded or failed

//HP := 0 haystack pointer
//NP := 0 needle pointer

for HP,_ := range haystack {
    //fmt.Println("start of outer loop")
    //fmt.Println("HP", HP, "NP", NP)
    start_of_matching = HP
    broken = false // fresh start of looking for needdle in the middle of haystack
    for NP,_ := range needle {
        //fmt.Println("start of inner loop")
        //fmt.Println("haystack value", haystack[HP],"needle value",needle[NP])
        if HP >= len(haystack) || haystack[HP] != needle[NP]  { // this char does not match OR haystack is longer than needle
            broken = true
            //fmt.Println("breaking", "HP", HP, "NP", NP)
            break
        } else { // this char matches}
            HP++ // 1 or more chars have matched, we need "match" the next char , NP++ is automatic, HP++ is not
            //fmt.Println("after HP++", "HP", HP, "NP", NP)
        }
        //fmt.Println("end of inner loop")
    } // FOR INNER
    if broken == false {
        ans = start_of_matching
        //fmt.Println("found answer")
        break
    }
    //fmt.Println("end of outer loop")
} // FOR OUTER
    
return ans

}

