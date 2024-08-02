package main

import (
	"fmt"
	"strings"
)

// variadatic function

func concatenate_all(s ...string) string {
	var cs string
	cs = strings.Join(s, " ")
	return cs
}

func main() {
	var result string
	result = concatenate_all("ab", "cd", "ef")
	fmt.Println(result)

}
