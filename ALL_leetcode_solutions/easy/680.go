// easy, 680, no ELO, ac rate 43pct
// first attempt
// my code worked feb 17  , but it was too slow
// performance matters that much for an "easy" problem?

// 2nd attempt: needed alg help, april 12th


func is_palin(s string) bool {
   l:=0
   r:=len(s)-1

   for l<r {
      if s[l] != s[r] {
        return false
      } else {
        l++
        r--
      }
   }

   return true
}



func validPalindrome(s string) bool {

  // testing func
  //s1 := "aaaaaabaaaaaa"
  //fmt.Println(is_palin(s1))


  l:=0
  r:=len(s)-1

  for l<r {
    if s[l] == s[r] {
        l++
        r--
    } else {
        temp1 := s[l+1:r+1] // remove left unmatched character
        temp2 := s[l:r] // remove right unmatched character
        //fmt.Println(temp1)
        //fmt.Println(temp2)
        if is_palin(temp1) || is_palin(temp2) {
            return true
        } else {
            return false
        }
    }
  }

    
  return true // didn't even have remove any , it was a palindrome to begin with
    
    
}
