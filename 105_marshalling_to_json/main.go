// MARSHALL STRUCT TO JSON

package main

import (
	"encoding/json"
	"fmt"
)

// Name and Rating MUST starts will Capital letters
type chess_player struct {
	Name   string
	Rating int
}

func main() {
	magnus := chess_player{Name: "magnus", Rating: 2800}

	resulting_json, err := json.Marshal(magnus)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(resulting_json))
	}

}
