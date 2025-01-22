// easy 1304 find uniq integers sum up to
// ELO 1167

// JAN 20, 2025

// problem lied when they said "any"
// answer must be line -2,-1,0,1,2

func sumZero(n int) []int {
    ans:=[]int{}

    if n == 1 {
        ans = append(ans,0)
        return ans
    }

    if n % 2 == 0 { // even
       low := (n/2) * (-1)
       high := n/2
       for i:=low;i<high+1;i++ {
        if i != 0 {
          ans = append(ans,i)
        }
      } // for
    }

    if n % 2 == 1 { // odd
      low := ((n-1)/2) * (-1) // e.g. if 3, then -1
      high := ((n-1)/2) // e.g. if 3, then 0
      for i:=low;i<high+1;i++ {
        ans = append(ans,i)
      } // for
    } // if

    return ans
    
} // func
