// leetcode easy #709 To Lower Case
// strings

/* My Test Case
"abcd"
"ABCD"
"Mamun"
"mamun"
"MAMUN"
"afhaklfhaKlll"
"HJHJKHKJahajdha"
"AbCdEfg"
"a&B"
*/

/* edge test case from leetcode
"al&phaBET"
*/




func toLowerCase(str string) string {
    
    answer := ""
    
    // To back and forth between "code" and "letter"
    // Tested in Go Playground
    // A-Z = 65-90
	// a-z = 97-122
    // diff = 32    (2 to the 5)
	//NUMCODE := rune('P')
	//fmt.Println(NUMCODE, string(NUMCODE))
	//LETTER := int(NUMCODE)
	//fmt.Println(LETTER)
    
    
    for _, char := range str {
        if rune(char) > 64 && rune(char) < 91 {
            // UPPER CASE
            new_numcode := rune(char) + 32 // takes us to lowercase
            answer = answer + string(new_numcode)
        } else { // not upper case , so leave it alone
            answer = answer + string(char)         
        }
    
    } // for loop
    
    return answer
}
