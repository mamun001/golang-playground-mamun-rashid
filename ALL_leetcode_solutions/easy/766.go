// easy (did not feel easy) 766 Toeplitz Matrix
// ELO 1249


func isToeplitzMatrix(matrix [][]int) bool {

    m := len(matrix)
    n := len(matrix[0])

    ans := true

    // you can have at most m+n-1 diagonals
    // i-j 
    // 6 diagonals
    // 4     0,3.        j-i = 3      i-j = -3 
    // 3,3   0,2 & 1,3.    j-i = 2.    i-j = -2
    // 2,2,2   0,1 & 1,2 & 2,3 j-i = 1.   i-j = -1
    // 1,1,1  0,0 & 1,1 & 2,2 j-i = 0.    i-j = 0
    // 5,5.    1,0 & 2,1 j-i = -1.        i-j = 1
    // 9       2,0 j-i=-2                  i-j = 2 


     // 3 diagonals
    // 2 0,1 i-j = -1
    // 1,2.   i-j = 0
    // 2 i-j = 1


    diag_index := 0
    // start at lower left i=m-1 j=0 i-j = diag_index = m-1
    // next : i=m-2 j=0 diag_index = m-2
    // last = top_right = i=0 j=m-1 diag_index = -m+1 =1-m
    // INTERVAL = i-j = diag_index goes down by 1 in each iteration

    map1 :=make(map[int]int)

    for i:=0; i<m;i++ {
        for j:=0 ; j<n;j++ {
            //fmt.Println(i,j,"value",matrix[i][j])
            diag_index = i-j
            //fmt.Println("diag_index",diag_index)
            _,exists := map1[diag_index]
            if exists {
                //fmt.Println("diag_index",diag_index," exists",value)
                if map1[diag_index] != matrix[i][j] {
                   return false
                }
            } else {
                map1[diag_index] = matrix[i][j]
            }
        }
    }


    return ans
    
}
