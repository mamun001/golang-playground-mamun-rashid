/*

leetcode easy #1119 remove vowels from string

my test cases

"ajhakhfwoiefwoiufwoiufwo"
"ajhfakiyewiowllalk"
"mhrvkytbgsraooy"
"bc"
"ae"
"ab"
"pw"
"ab"
"ba"
"a"
"b"

*/





func removeVowels(s string) string {
    answer := ""
    
    for _, char := range s {
        if char != 'a' && char != 'e' && char != 'i' && char != 'o' && char != 'u' {
            answer = answer + string(char)
        }
    }
    
    return answer
    
}
