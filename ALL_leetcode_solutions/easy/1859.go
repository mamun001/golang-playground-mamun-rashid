// easy, 1859, ELO 1290
// sorting the sentence

import "strings"

func char(n int) byte {
    switch n {
        case 0: return '0'
        case 1: return '1'
        case 2: return '2'
        case 3: return '3'
        case 4: return '4'
        case 5: return '5'
        case 6: return '6'
        case 7: return '7'
        case 8: return '8'
        case 9: return '9'
    }
    return '0'
}

func sortSentence(s string) string {

   words:=strings.Split(s, " ")
   //fmt.Println(words)

   ans:=""

   //fmt.Println(byte('1'))  '1' = 49
   //fmt.Println(char(5))

   indexed_words:=make(map[int]string) // we will like 1 = firstword

   for i:=0;i<=len(words);i++{
     //fmt.Println(i,string(i))
     for j:=0;j<len(words);j++{
      //fmt.Println(words[j][len(words[j])-1])
       if words[j][len(words[j])-1] == char(i) { // if this word ends with i (1,2,3 etc)
          //fmt.Println(words[j])
          indexed_words[i]=words[j] // insert into the map , this word with the i as the key
       }
     }
   }

   //fmt.Println(indexed_words)
   // now we have a map of words like 1 = firstword1, 2=secondword2

   // put together the final sentence by removing the integer from the end of each word and adding space
   for i:=1;i<=len(words);i++{
      ans=ans+indexed_words[i][:len(indexed_words[i])-1]+" "
   }

   //take out the last space
   ans=ans[:len(ans)-1]

    return ans


    
}
