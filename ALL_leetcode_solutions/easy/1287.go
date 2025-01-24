// easy 1287, ELO 1179
// elemnet appearing more than 25%

// Jan 22, 2025
// my alg perf: cpu beats 100%, RAM beats 76%

func findSpecialInteger(arr []int) int {

    ans := -1

    target_count := (len(arr) / 4) -1 // sometimes the count can fall right on the border

    if len(arr) == 1 {
        return arr[0]
    }

    count := 0 // how of this number has come so far

    for i:=0 ; i< len(arr) ; i++ {
        if i == 0 {
            count = 1
        }
        if i > 0 && arr[i] == arr[i-1] { // same number continues 
            count++
            if count > target_count {
             // found our answer
             return arr[i]
           }
        } else {  // brand new number
            count = 0
        } // else
    } // for

    return ans
    
} // func
