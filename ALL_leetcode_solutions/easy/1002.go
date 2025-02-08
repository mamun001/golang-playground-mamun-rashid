// easy 1002, ELO 1280
// find common characters

func commonChars(words []string) []string {


    // 97 = a
    // 98 = b
    // 99 = c
    // 100 = d
    // 101 = e
    // 108 = l
    // 111 = o

    ans :=[]string{}

    first_map := make(map[rune]int) // save frequencies of letters in the first word
    this_map := make(map[rune]int) // save frequencies of the word being looked at

    // save frequencies of letters in the first word
    for i,_ := range words[0] {
        first_map[rune(words[0][i])]++
    }
    //fmt.Println("first_map initial",first_map)



  // LOOP THROUGH EACH WORD
  for i:=1;i<len(words);i++ {

    // zero out this word's letter frequency map
    for k,_:= range this_map {
        this_map[k] = 0
    }


    // build this word's letter frequency map
    for j,_ := range words[i] {
        this_map[rune(words[i][j])]++
    }
    //fmt.Println(" ")
    //fmt.Println(words[i])
    //fmt.Println("i and this_map",i,this_map)


    // MERGE THE TWO MAPS such that we keep lower frequency for a given letter
    for k,v:= range this_map {
        if first_map[k] > this_map[k] {
            first_map[k] = v
        }
    }
    for k,_:= range first_map {
        if this_map[k] == 0 {
            first_map[k] = 0
        }
    }
    //fmt.Println("first_map after i is done merging two maps",i,first_map)
  }




  //fmt.Println("_____end_result_________")
  //fmt.Println("first_map",first_map)


   // build one letter strings form the end result map
   // none, if zero
   // 1 letter string if 1
   // N 1 letter strings in more than 1
   for k,v:= range first_map {
       if v > 0 {
          //fmt.Println(string(k))
          for i:=0;i<v;i++ {
             ans = append(ans,string(k))
          }
       }
   }





    return ans


    
}
