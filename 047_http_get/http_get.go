package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("http://www.example.com")

	if err != nil {
		fmt.Print("There was an error from URL")
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Print("There was an error reading incoming stream")
	}

	fmt.Print(body)

}
