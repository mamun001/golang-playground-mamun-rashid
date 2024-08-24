package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "T is M hajkhklqhkqj"

	longest_word := ""
	longest_word_length := 0

	words := strings.Split(s1, " ")
	//fmt.Println(words)

	for i, word := range words {
		if len(word) > longest_word_length {
			longest_word = words[i]
			longest_word_length = len(word)
		}
		//fmt.Println(i, word, longest_word, longest_word_length)
	}

	fmt.Println(longest_word)

}
