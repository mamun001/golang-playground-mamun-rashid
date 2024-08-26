// UNMARSHALL JSON TO STRUCT

package main

import (
	"encoding/json"
	"fmt"
)

type chess_player struct {
	Name   string
	Rating int
}

func main() {
	// json_data is string that has the fields embedded in it
	json_data := `{"name": "magnus", "rating": 2800}`

	// making an empty object that will eventually hold the "struct" data, once unmarshalled
	var magnus chess_player

	// unmarshalling
	err := json.Unmarshal([]byte(json_data), &magnus)

	if err != nil {
		fmt.Println("error occurred", err)
	} else {
		fmt.Println(magnus)
	}

}
