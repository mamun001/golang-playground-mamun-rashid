package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	f, err := os.Open("foo.log")

	if err != nil {
		fmt.Println(err)
	}

	content, err := io.ReadAll(f)

	defer f.Close()

	fmt.Println(string(content))

}
