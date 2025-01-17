// 73 pct acceptance rate

// 2351 First Letter to appear twice

// easy level

// DEC 26, 2024

import "fmt"

func repeatedCharacter(s string) byte {

    keeping_track := make(map[rune] int)
    ans := 'Z' // default answer denoting we did not find a duplicate

    for _,v := range s {
       //fmt.Println(byte(v))
       
       keeping_track[v]++
       if keeping_track[v] == 2 {
          ans = v
          break
       }


    }
    
    //fmt.Println(keeping_track)

    return byte(ans)

}

