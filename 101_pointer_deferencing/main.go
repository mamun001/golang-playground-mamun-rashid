package main

import "fmt"

// 123: Change a struct's value inside a func
// Equivalent to Pointer Dereferencing
// example: (*chessplayer1).rating = 2200

type ChessPlayer struct {
	name   string
	rating int
}

func UpRatingBy100(chessplayer1 *ChessPlayer) {
	// chessplayer1.rating = chessplayer1.rating + 100
	(*chessplayer1).rating = (*chessplayer1).rating + 100

}

func main() {

	John_Doe := ChessPlayer{
		name:   "John Doe",
		rating: 2100}

	fmt.Println(John_Doe)
	UpRatingBy100(&John_Doe)
	fmt.Println(John_Doe)

}
