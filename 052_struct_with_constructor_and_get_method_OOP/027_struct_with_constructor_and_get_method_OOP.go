package main

// Tested on Go Playground

import (
	"fmt"
)

type book struct {
	id           int
	title        string
	author       string
	publish_year int
}

// constructor
func newBook(id int, title string, author string, publish_year int) *book {

	b := book{id: id, title: title, author: author, publish_year: publish_year}
	return &b
}

func getBookTitle(onebook *book) string {
	return onebook.title
}

//

func main() {
	book1 := newBook(101, "Foo Title", "john doe", 1999)
	fmt.Println(book1.id)
	fmt.Println(book1.title)
	fmt.Println(book1.author)
	fmt.Println(book1.publish_year)

	// get the title of a book
	fmt.Println(getBookTitle(book1))
}

