// easy 2133, ELO 1264
// check if every row and column contains all numbers

func is_in(ar []int, n int) bool {
    for i:=0;i<len(ar);i++ {
        if ar[i] == n {
            return true
        }
    }
    return false
}

func is_valid_slice(ar []int) bool {
    N:=len(ar)
    for i:=1;i<=N;i++ {
        if is_in(ar,i) == false {
            return false
        }
    }
    return true
}

func checkValid(matrix [][]int) bool {

    //fmt.Println(is_valid_slice(matrix[3]))
    L:=len(matrix)

    for i:=0;i<L;i++ {
      column:=[]int{}
      for j:=0;j<L;j++ {
          column=append(column,matrix[j][i])
      }
      //fmt.Println(column)
      if is_valid_slice(column) == false {
          return false
      }
    }
    
    for i:=0;i<L;i++{
        if is_valid_slice(matrix[i]) == false {
            return false
        }
    }

    

    return true
    
}
