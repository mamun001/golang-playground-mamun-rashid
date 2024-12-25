// 2942 Find Words Containing Character

// acc rate 88.5%

// DEC 24, 2024


func findWordsContaining(words []string, x byte) []int {

    ans := []int{}

    for i,_ := range words { // loop over each word in array
        for _,z := range words[i] { // loop over each character in a given word
            if z == rune(x) {
                ans = append(ans,i)
                break // do not repeat IF the same word has the given character more than once
            }
        }
    }

    return ans
}
