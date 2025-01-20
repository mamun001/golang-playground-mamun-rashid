// 3274 easy , check if two
// ELO 1162

// JAN 18, 2025

import "fmt"


func abs(n int) int {
    ans := n
    if n < 0 {
        ans = n * -1
    }
    return ans
}



func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {

    
  ans := false

  //fmt.Println('a' - 'b')
  //fmt.Println(coordinate1[0])
  diff1 := abs(int(coordinate2[0]-coordinate1[0]))
  diff2 := abs(int(coordinate2[1]-coordinate1[1]))
  sum := diff1 + diff2
  //fmt.Println(diff1)
  //fmt.Println(diff2)

  if sum % 2 == 0 {
    ans = true
  }

  return ans

}
