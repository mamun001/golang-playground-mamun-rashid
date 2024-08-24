

// leetcode number 1773 easy
//
// Mamun's code
//
// nested loops
//


func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    
    answer := 0
    second_index := 99
    
    if ruleKey == "type" {
        second_index = 0
    }
    
    if ruleKey == "color" {
        second_index = 1
    }
    
    if ruleKey == "name" {
        second_index = 2
    }
    
    
    
    for i:=0; i<len(items);i++ {
        //fmt.Println(items[i][0])
        for j:=0;j<3;j++ {
          //fmt.Println(items[i][j])
            if second_index == j && items[i][j] == ruleValue {
                answer++
            }
        
        }    
    }
    
    return answer
    
    
}
