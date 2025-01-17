import (
    "strconv"
)


func isPalindrome(x int) bool {
    answer := true // start assuming it is palindrome

    if x < 0 {
        answer = false // negative number is always NOT a palindrome
    }
   

    if x == 0 {
        answer = true // 0 is special case or maybe , so get it out of the way
    }

     

    if x > 0 {
      string1 := strconv.Itoa(x)
      len1 := len(string1)
      // loop through each character and compare with its equivalent from the end
      for i :=0 ; i < len1 ; i++ {
        if string1[i] != string1[len1-i-1] {
            answer = false // no match, answer is false and remain false through the loop
        }
       
      }
    }
   

    return answer
}
