
package main


// Test in go playground
import (
	"fmt"
)

func main() {

	// A-Z = 65-90
	// a-z = 97-122
        // differnce between Upper Case and Lower Case 32
	NUMCODE := rune('P')
	fmt.Println(NUMCODE, string(NUMCODE))
	LETTER := int(NUMCODE)
	fmt.Println(LETTER)
}

