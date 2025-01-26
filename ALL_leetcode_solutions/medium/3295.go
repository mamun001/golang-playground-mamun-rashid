// MEDIUM 3295 ELO 1198 report spam messages
// my alg worked BUT failed on scale 100,000
// 
// KEY THING: IF you make a map of words and do a lookup againts INSREAD OF
//.  CHECKING YOURSELF : LINEAR vs O of MxN
// That's what makes the difference


// my original lookup by hand O of MxN
//func is_in (ar []string, s string) bool {
    //ans := false
    //for i:=0; i <len(ar) ; i++ {
        //if s == ar[i] {
          //return true
       // }
    //}
    //return ans
//}


func reportSpam(message []string, bannedWords []string) bool {

  ans := false
  found_count := 0
  found_two := false

  // THEY KEY, make a map and look up agaist it. LINEAR!!
  map_b := make(map[string]int)
  
  for i:=0 ; i < len(bannedWords); i++ {
    map_b[bannedWords[i]] = 1
  }
  

  for i:=0 ; found_two == false && i < len(message) ; i++{
       if map_b[message[i]] == 1 {
          found_count++
          if found_count > 1 {
            return true
            break
          }
       }
       
   }

    return ans
    
}

