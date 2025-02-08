// easy 2678, ELO 1198
// number of senior citizens

import "strconv"

func countSeniors(details []string) int {

  age := ""
  age2 := 0
  ans := 0

   for i:=0; i<len(details); i++ {
      age = details[i][11:13]
      //fmt.Println(age)
      age2,_ = strconv.Atoi(age)
      //fmt.Println(age)
      if age2 > 60 {
        ans++
      }
   }

   age2 = age2 // just so golang will not complain

   return ans
    
}

