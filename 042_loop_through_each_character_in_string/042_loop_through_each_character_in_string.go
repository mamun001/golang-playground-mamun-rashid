package main

// Tested in golang playground
// 
// Mamun's code
//

import (
	"fmt"
)

func main() {

        // rune_value is basiclaly is character (turns out to be of type rune)
        // since rune_value holds one "letter" at a time, you can now do what you want with each letter
	for _, rune_value := range "mamunrashid" {
		fmt.Println(string(rune_value))
	}

}

