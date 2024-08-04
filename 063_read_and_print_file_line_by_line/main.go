package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var filename string = "logs.txt"

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Could not open file", filename)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}
