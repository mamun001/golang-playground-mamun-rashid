// easy 3340 , check balanced string, ELO 1190
// JAN 24, 2025


func to_i (b byte) int {
    ans := 0
    switch b {
        case '0':
          return 0
        case '1':
          return 1
        case '2':
          return 2
        case '3':
          return 3
        case '4':
          return 4
        case '5':
          return 5
        case '6':
          return 6
        case '7':
          return 7
        case '8':
          return 8
        case '9':
          return 9
    }
    return ans
}

func isBalanced(num string) bool {

   // fmt.Println(to_i(num[0]))

   sum_o := 0
   sum_e := 0
   
   for i:=0 ; i < len(num); i ++ {
       if i % 2 == 0 {
           sum_e = sum_e + to_i(num[i])
       } else {
           sum_o = sum_o + to_i(num[i])
       }
   }

   //fmt.Println(sum_o,sum_e)

    return sum_o == sum_e

    
}
