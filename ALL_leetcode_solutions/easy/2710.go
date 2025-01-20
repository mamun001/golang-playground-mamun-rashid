// ELO 1164
/ 2710 Remove Trailing Zeros

// Jan 18, 2025

func removeTrailingZeros(num string) string {
 
    last_non_zero := 0

    for i,v := range num {
        if v != '0' {
            last_non_zero = i
        }

    }


    return num[:last_non_zero+1]
    
}
