package main

import "fmt"

type CP struct {
	Name   string
	Rating int
}

func (cp1 *CP) PrintName() {
	fmt.Println(cp1.Name)
}

func main() {

	John_Doe := CP{Name: "John Doe", Rating: 2100}

	pn := John_Doe.PrintName

	pn()
	pn()
	pn()
	pn()

}
