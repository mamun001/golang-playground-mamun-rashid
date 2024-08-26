package main

import "fmt"

type chess_player struct {
	name           string
	rating         int
	played_against []string
}

func main() {
	magnus := chess_player{
		name:           "magnus",
		rating:         2800,
		played_against: []string{"hikaru", "nepo"},
	}

	fmt.Println(magnus)
}
