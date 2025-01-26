//easy 1880 check if word equals 
// ELO 1187

// Jan 24, 2025


// integer value of char
func val (r rune) int {
    // lazy coding
    ans := 0

    if r == 'a' {
       return 0
    } else if r == 'b' {
        return 1
    } else if r == 'c' {
        return 2
    } else if r == 'd' {
        return 3
    } else if r == 'e' {
        return 4
    } else if r == 'f' {
        return 5
    } else if r == 'g' {
        return 6
    } else if r == 'h' {
        return 7
    } else if r == 'i' {
        return 8
    } else if r == 'j' {
        return 9
    }

    return ans
}

// power of 10
func po(n int) int {
    ans := 1
    for i:=0; i < n; i++ {
        ans = ans * 10
    }
    return ans
}

// main func
func isSumEqual(firstWord string, secondWord string, targetWord string) bool {

    int_val_1 :=0 // value of firstWord
    n := len(firstWord)
    for i:=0 ; i<n; i++ {
        int_val_1 = int_val_1 + val(rune(firstWord[i])) * po(n-i-1)
        //fmt.Println(i,n-i,int_val)
    }
    //fmt.Println(int_val_1)

    int_val_2 :=0 // value of secondWord
    n = len(secondWord)
    for i:=0 ; i<n; i++ {
        int_val_2 = int_val_2 + val(rune(secondWord[i])) * po(n-i-1)
        //fmt.Println(i,n-i,int_val)
    }
    //fmt.Println(int_val_2)

    int_val_3 :=0 // value of targetWord
    n = len(targetWord)
    for i:=0 ; i<n; i++ {
        int_val_3 = int_val_3 + val(rune(targetWord[i])) * po(n-i-1)
        //fmt.Println(i,n-i,int_val)
    }
    //fmt.Println(int_val_3)

    return int_val_1 + int_val_2 == int_val_3
    
}

