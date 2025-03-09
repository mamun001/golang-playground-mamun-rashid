// easy 2716, ELO 1243
// all that mumbo jumbo in the probelm description = count the uniq characters in the string

func minimizedStringLength(s string) int {

    map1 := make(map[byte]int)

   for i:=0;i<len(s);i++ {
       map1[s[i]]++
   }

    return len(map1)
    
}
