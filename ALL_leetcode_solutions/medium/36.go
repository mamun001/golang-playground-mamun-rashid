// MEDIUM, 36, success without help after 3 hours of coding and debugging
// part of Neetcode 250
// no ELO

// first 2/3 of code is horrendous because by the time (last one third , checking 3x3), I figured out
// that I shoudl simply create a func to check an array of 9 elements, I had already done the first 2/3 
// (row and column checks), and I did not feel like refactoring

func is_correct(ar []byte) bool {
    ans := true
    map1 := make(map[byte] int)

    for i:=0;i<len(ar);i++ {
        if ar[i] != '.' {
              map1[ar[i]]++
           }
           if map1[ar[i]] > 1 {
               return false
           }

    }

    return ans

}

func reset(m map[byte] int) map[byte] int {
    ans := make(map[byte] int)
    ans['1'] = 0
    ans['2'] = 0
    ans['3'] = 0
    ans['4'] = 0
    ans['5'] = 0
    ans['6'] = 0
    ans['7'] = 0
    ans['8'] = 0
    ans['9'] = 0
    return ans
}

func isValidSudoku(board [][]byte) bool {

    

    // . = 46
    // 1 = 49
    // 2 = 50
    // 3 = 51 
    // 4 = 52
    // 5 = 53
    // 8 = 56
    // 9 = 57

    ans := true

    map1 := make(map[byte] int)
    reset(map1)

    //testing func is_correct
    //ar := []byte{'1','2','3','4','5','6','7','9','9'}
    //fmt.Println("is correct text",is_correct(ar))

    // check each row
    for i:=0;i<9;i++{
       //fmt.Println(i)
       // reset the map
       for k:=0;k<9;k++{
          map1 = reset(map1)
       }
       //fmt.Println(map1)
       for j:=0;j<9;j++{
           //fmt.Println(j)
           if board[i][j] != '.' {
              map1[board[i][j]]++
           }
           if map1[board[i][j]] > 1 {
               //ans = false
               //fmt.Println("false row",i,j,map1)
               return false
           }
       }
       //fmt.Println(map1)
    } 

    //fmt.Println("row to column")

    
    // check each column
    for j:=0;j<9;j++{
       //fmt.Println("j",j)
       // reset the map
       for k:=0;k<9;k++{
          map1 = reset(map1)
       }
       //fmt.Println("map should be reset",map1)
       for i:=0;i<9;i++{
           //fmt.Println("i",i)
           if board[i][j] != '.' {
              map1[board[i][j]]++
              //fmt.Println(map1)
           }
           if map1[board[i][j]] > 1 {
               //ans = false
               //fmt.Println("false column",i,j,map1)
               return false
           }
       }
       //fmt.Println(map1)
    } 



    // check the 3x3s

    for y:=1;y<9;y=y+3{
        for z:=1;z<9;z=z+3{
            //fmt.Println(m,n)
            ar := []byte{}
            for i:=y-1;i<y+2;i++{
                for j:=z-1;j<z+2;j++{
                   ar = append(ar,board[i][j])
                }
             }
            //fmt.Println(ar)
            if is_correct(ar) == false {
                //ans = false 
                //fmt.Println("false 3x3 y z ar",y,z,ar)
                return false
            }
        }
    }

    
            
    return ans
    
}
