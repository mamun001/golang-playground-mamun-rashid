
/* 
#1832 Pangram
*/


/*
test input data I made up
"thequickbrownfoxjumpsoverthelazydog"
"leetcode"
"abcdefghijklmnopqrstuvwxyz"
"abc"
"abcdefghjklmnopqrtsuvwxy"
"bcdefghjklmnopqrtsuvwxyz"
"zyxwuvutsrqpmnokljhbcdefgh"
"zyxwuvutsrqpmnokljhb"
*/


func checkIfPangram(sentence string) bool {
    
    answer := true // we start with assumption that senetence IS a pangram
    
    map1 := make(map[int]bool)   
    
    // 97=a 122=z
    // fill up map1 with a to z = false 
    for i:=97; i <= 122; i++ {
        map1[i]=false
    }
    
    //fmt.Println(map1)
    
    // whatever we find in "sentence" , we mark key in map1 to be "true"
    key := 99999 // init the value with insane number
    for _, char := range sentence {
        key=int(char) // e.g. if char is then key = 97
        map1[key] = true
    }
    
    
    fmt.Println("----------------------------")
    fmt.Println(map1)
    fmt.Println("----------------------------")
    
    // now loop through map1 looking for false = that will make the incoming string NOT a pangram
    
    
    for _, value := range map1 {
        if value == false {
            answer = false
            break
        }
    }
    

    
    return answer
}
