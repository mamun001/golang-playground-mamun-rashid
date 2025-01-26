// reverse an integer
//
package main

import (
"fmt"
"strconv"
)

func main() {

i := 1234567
s := strconv.Itoa(i)
rs := "" // reversed string
N := len(s)
for i := N - 1; i > -1; i-- {
rs = rs + string(s[i])
}
ri, _ := strconv.Atoi(rs) // reversed integer
fmt.Println(s) // 1234567
fmt.Println(rs) // 7654321
fmt.Println(ri) // 7654321
}

