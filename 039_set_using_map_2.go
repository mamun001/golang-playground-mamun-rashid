package main

import "fmt"

func main() {

chess_pieces := make(map[string]struct{})

chess_pieces["pawn"] = struct{}{}  // insert
chess_pieces["king"] = struct{}{}  // insert
chess_pieces["queen"] = struct{}{} // insert
fmt.Println(chess_pieces)

delete(chess_pieces, "queen") // delete using built-in delete func of maps
fmt.Println(chess_pieces)

_, exist_var := chess_pieces["pawn"] // does "pawn" exist? again, using map's built-in functionality
fmt.Println(exist_var)

_, exist_var = chess_pieces["knight"] // does "pawn" exist? again, using map's built-in functionality
fmt.Println(exist_var)
}

