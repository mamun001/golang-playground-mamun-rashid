package main

import "fmt"

// EMBEDDING

type ChessPlayer struct {
	Name   string
	Rating int
}

type GrandMaster struct {
	ChessPlayer
	GMYear int
}

func PrintRating(cp ChessPlayer) {
	fmt.Println(cp.Rating)
}

func main() {

	John_Doe := GrandMaster{
		ChessPlayer: ChessPlayer{
			Name:   "John Doe",
			Rating: 2500,
		},
		GMYear: 1999,
	}

	fmt.Println(John_Doe.ChessPlayer)
	fmt.Println(John_Doe.ChessPlayer.Name)
	fmt.Println(John_Doe.GMYear)

	PrintRating(John_Doe.ChessPlayer)

}
