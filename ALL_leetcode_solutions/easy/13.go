

// THIS BEAT 100% of the submissions in SPEEED! 12.19.2024

func romanToInt(s string) int {
    sum := 0
    for i := 0 ; i < len(s) ; i++ {
        switch s[i] {
            case 'I':
              sum = sum + 1

            case 'V':
              if i != 0 && s[i-1] == 'I' {
                  sum = sum +3 // minus 1 for previous I and minus 1 for because this is 4 not 5
              } else {
                  sum = sum + 5
              }

            case 'X':
              //
             if i != 0 && s[i-1] == 'I' {
                  sum = sum + 8 // minus 1 for previous I and minus 1 for because this is 9 not 10
              } else {
                  sum = sum + 10
              }

            case 'L':
              if i != 0 && s[i-1] == 'X' {
                   sum = sum + 30 // minus 10 for previous X and minus 10 for because this is 40 not 50
              } else {
                  sum = sum + 50
              }

            case 'C':
              //
              if i != 0 && s[i-1] == 'X' {
                   sum = sum + 80 // same reasons
              } else {
                  sum = sum + 100
              }

            case 'D':
              if i != 0 && s[i-1] == 'C' {
                  sum = sum + 300 // same reasons as above
              } else {
                  sum = sum + 500
              }

            case 'M':
              //
              if i != 0 && s[i-1] == 'C' {
                  sum = sum + 800 // same reasons as above
              } else {
                  sum = sum + 1000
              }

        }
    }
    

    return sum
}
