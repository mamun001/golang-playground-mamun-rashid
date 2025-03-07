// easy 1221 split a string in balanced strings
// JAN 14, 2025

// I swear, code worked in 1 shot, not even a syntax error

// my alg: keep count Ls and Rs , whenever Ls = Rs , we have a substring

func balancedStringSplit(s string) int {

    Lsofar := 0
    Rsofar := 0
    BST := 0 // balanced substrings count so far, because at 0 the following code will find a balanced substring

    for i:=0 ; i<len(s); i++ {
        if s[i] == 'L' {
            Lsofar++
        } else {
            Rsofar++
        }
        
        if Lsofar == Rsofar {
            BST++
        }
    }

    return BST
    
}

