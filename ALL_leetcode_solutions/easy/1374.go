// 1374 generate a string
// ELO 1164 or so

// JAN 18, 2025

func generateTheString(n int) string {

    ans := ""

    if n % 2 == 1 { // odd
       for i:=0 ; i<n; i++ {
          ans = ans + string('a')
       }
    }

    if n % 2 == 0 { // even
       for i:=0 ; i<n-1; i++ {
          ans = ans + string('a')
       }
       ans = ans + string('b')
    }
    

   return ans
    
}

