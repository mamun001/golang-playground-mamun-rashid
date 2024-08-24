package main

/*

1. Contains
2. Fields
3. Index
4. Clone
5. Compare
6. Join
7. Count
8. Repeat
9. Replace
10. Split
11. ToLower
12. ToUpper
13. Trim
14. TrimSpace
15. HasPrefix

*/

import (
	"fmt"
	"strings"
)

func main() {
	var s1 string
	var s2 string

	s1 = "      Abc Def ghg jjj kkkk      "

	s2 = strings.TrimSpace(s1)

	fmt.Println(s2)
}
