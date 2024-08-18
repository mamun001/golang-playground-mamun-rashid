package main

import (
	"fmt"
)

type Chess_Player struct {
	Name   string
	Rating int
}

func up_rating_by_100(cp1 *Chess_Player) {
	cp1.Rating = cp1.Rating + 100
}

func main() {
	John_Doe := Chess_Player{Name: "John Doe", Rating: 2100}
	fmt.Println(John_Doe)

	up_rating_by_100(&John_Doe)
	fmt.Println(John_Doe)

}
