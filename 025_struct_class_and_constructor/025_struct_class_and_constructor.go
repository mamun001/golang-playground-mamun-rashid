
// Looks like Mamun's code
//
// Tested on Go Playground
//
// TASK: use struct to make BOOK and write FUNC for NEWBOOK CONSTRUCTOR


package main
import (
	"fmt"
)

type book struct {
	title        string
	author       string
	publish_year int
}

// constructor
func newBook(title string, author string, publish_year int) *book {

	b := book{title: title, author: author, publish_year: publish_year}
	return &b
}

func main() {
	book1 := newBook("Foo Title", "john doe", 1999)
	fmt.Println(book1.title)
	fmt.Println(book1.author)
	fmt.Println(book1.publish_year)
}

