// easy 2194, ELO 1253
// Cells in a Range on an Excel Sheet

func cellsInRange(s string) []string {

  ans:=[]string{}

  start_col := byte(s[0])
  end_col := byte(s[3])
  //fmt.Println(start_col,end_col)
  start_row := byte(s[1])
  end_row := byte(s[4])
  //fmt.Println(start_row,end_row)

  for i:=start_col;i<=end_col;i++ {
    for j:=start_row;j<=end_row;j++ {
        //fmt.Println(i,j)
        temp:=string(i)+string(j)
        ans=append(ans,temp)
        //fmt.Println(ans)
    }
  }

  return ans 
}
