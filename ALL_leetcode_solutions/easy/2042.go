// easy 2042, ELO 1257
// check if the numbers are ascending in a sentence

import "strconv"


// check if the latest number is larger than last number or not
func check_the_number(s string, last_number int) bool {
    num,_:=strconv.Atoi(s)
    if num <= last_number {
        return false
    }
    return true
}

func areNumbersAscending(s string) bool {

   this_number_string:=""
   last_number:=0
   
   for i:=0;i<len(s);i++ {
      if s[i] >= '0' && s[i] <= '9' { // digit
          this_number_string=this_number_string+string(s[i])
          //fmt.Println(this_number_string)
      } else if s[i] == ' ' && s[i-1] >= '0' && s[i-1] <= '9' { // number ended
          //fmt.Println(this_number_string)
          //fmt.Println(check_the_number(this_number_string,last_number))
          if check_the_number(this_number_string,last_number) == false {
            //fmt.Println("returning false")
            return false
          } else {
            //fmt.Println("check number func returned true")
            last_number,_=strconv.Atoi(this_number_string)
            this_number_string=""
          }
      }
   }

   if s[len(s)-1] >= '0' && s[len(s)-1] <= '9' {
     //fmt.Println(this_number_string)
     if check_the_number(this_number_string,last_number) == false {
            //fmt.Println("returning false")
            return false
     }
   }

   return true
    
}
