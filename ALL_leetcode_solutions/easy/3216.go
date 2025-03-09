// easy, 3216, ELO 1243
// lexigraphically smallest string after a swap

// all that mumbo jump in the problem , basically means
// find first pair of odd or even digits next together where 2nd one is smaller than the first and swap the two


func is_in(ar []byte, b byte) bool {
    for i:=0;i<len(ar);i++ {
        if ar[i] == b {
            return true
        }
    }
    return false
}

func same_parity(b1,b2 byte) bool {
    odds:=[]byte{'1','3','5','7','9'}
    evens:=[]byte{'0','2','4','6','8'}

    if is_in(odds,b1) && is_in(odds,b2) {
        return true
    }

    if is_in(evens,b1) && is_in(evens,b2) {
        return true
    }

    return false


}

func getSmallestString(s string) string {

    //fmt.Println(same_parity(s[1],s[2]))

    // '9' > '0'
    //fmt.Println('9'>'0')

    // loop through and find the first such pair and make a new string by swapping them
    for i:=0;i<len(s)-1;i++ {
        if same_parity(s[i],s[i+1]) && s[i+1] < s[i] {
            return s[:i]+string(s[i+1])+string(s[i])+s[i+2:]
        }
    }


    return s
    
}
