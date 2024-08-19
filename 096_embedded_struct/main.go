package main

import "fmt"

type ChessPlayer struct {
	Name   string
	Rating int
}

type GrandMaster struct {
	ChessPlayer
	GMYear int
}

func main() {

	John_Doe := GrandMaster{
		ChessPlayer: ChessPlayer{
			Name:   "John Doe",
			Rating: 2100},
		GMYear: 1990,
	}

	fmt.Println(John_Doe.Rating)
}
