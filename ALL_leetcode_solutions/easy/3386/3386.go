// easy 3386
// acc rate 40.6
// button with longest push time

func buttonWithLongestTime(events [][]int) int {

    current_ans := 10001 // start of value
    longest_time := 0
    time_diff := 0
    

    for i := 0 ; i < len(events) ; i++ {
        if i == 0 {
            time_diff = events[i][1]
        } else {
            time_diff = events[i][1] - events[i-1][1]
        }
        if time_diff == longest_time { // tie
            longest_time = time_diff
            if events[i][0] < current_ans { // events[i][0] is latest high water mark, ans was the previous water mark, lower one wins tie
                current_ans = events[i][0] // new winner is smaller button number
            }
        }
        if time_diff > longest_time { // outright winner
            longest_time = time_diff
            current_ans = events[i][0]
            }
        } // for loop
        return current_ans
    } // func

