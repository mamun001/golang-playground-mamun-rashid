

// leetcode easy #1323 Maximum 69 number
// math and strings


/* my data set, max 10000 by definition
9999
9996
6
9
69
96
999
996
9969
9996
*/



func maximum69Number (num int) int {
    
    // from stackoverflow, convert int to string
    s :=strconv.Itoa(num)
    
    //fmt.Println("original:",s)
    
    new_string := ""
    switched_already := false
    
    
    // loop through each digit and build out "new_string" with switching first 6 that you see
    for _, char := range s {
        //fmt.Printf("character %c starts at byte position %d\n", char, pos)
        
        if switched_already == false && string(char) == "6"{
            //fmt.Println("switching") 
            new_string = new_string + string('9') // switching
            switched_already = true
        } else {
            new_string = new_string + string(char)  // no switching needed
        }
            
    } // for
        
    
    //fmt.Println("after switching:",new_s)
    // from stackoverflow
    x, _ := strconv.ParseInt(new_string, 10, 64)
    
    return int(x)
    
    
}
