// easy 3238 ELO 1285
// Find the number of winnning


func colors(ar [][]int, n int) int {
    // how many colors did n win of the same color, whichever that is?
    max:= 0
    M :=len(ar)
    //N :=len(ar[0])
    map1:=make(map[int]int) // colors and counts

    //fmt.Println(M,N,n)

    // loop over arr (2D), but Mx2
    for i:=0;i<M;i++{
       if ar[i][0] == n { // if we are talking about the player fed to this func
          map1[ar[i][1]]++ // 2nd index in each element of ar has the color code
       }
    }
    //fmt.Println(map1)

    for _,v := range map1 {
        if v > max {
            max = v
        }
    }

    return max
}

func winningPlayerCount(n int, pick [][]int) int {

  ans := 0

  //fmt.Println(colors(pick,5))
  for i:=0;i<n;i++ {
    if colors(pick,i) > i {
        ans++
    }
  }


  return ans
    
}
