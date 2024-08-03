package main

import (
	"fmt"
	"strings"
)

// see the data set below
//
// Now write a function that will take in something that will query this array
// for particular value for the first word in the 2nd column
// (i.e. batter or bowler etc.) and same for the 2nd and 3rd.
// That is to say, one can call this function from main and give it specific values
// for each these categories and get the first names ONLY that match all those categories supplied
// categories can be 0 to 3.

// Note that in the 2nd column , the order of the data may vary. They will not always come in the same order.

// write your code below
// you will have to change the func defition for sure

func get_names(data [][2]string, match_these ...string) []string {
	var result []string
	var do_not_add_result bool

	for index_data := range data {
		do_not_add_result = false
		for index_match_these := range match_these {
			if strings.Contains(data[index_data][1], match_these[index_match_these]) == false {
				do_not_add_result = true
			}
		}

		if do_not_add_result == false {
			result = append(result, data[index_data][0])
		}
	}

	return result
}

func main() {
	// Create a 10x2 slice of strings
	data := make([][2]string, 0)

	// Assign the specified values
	data = append(data, [2]string{"Alice", "catcher,dodgers,highpay"})
	data = append(data, [2]string{"Bob", "bowler,giants,freeagent"})
	data = append(data, [2]string{"Charlie", "batter,angels,freeagent"})
	data = append(data, [2]string{"David", "batter,giants,highpay"})
	data = append(data, [2]string{"Eve", "batter,angels,highpay"})
	data = append(data, [2]string{"Frank", "bowler,dodgers,freeagent"})
	data = append(data, [2]string{"Grace", "catcher,dodgers,highpay"})
	data = append(data, [2]string{"Hannah", "bowler,angels,highpay"})
	data = append(data, [2]string{"Ivy", "batter,giants,lowpay"})
	data = append(data, [2]string{"Jack", "bowler,dodgers,highpay"})

	fmt.Println(get_names(data, "giants", "lowpay", "batter"))

}
