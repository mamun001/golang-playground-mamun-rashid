// easy 989, ELO 1235
// Add to array-Form of integer
// Convert to int --> add --> convert back to array fails (testes!), beacuse test cases includes numbers
// that goes beyond int32 range!

// HENCE THIS alg: convert k to array , add digit by digits
// However, I should have just defined ans to be max(len of num and k-digits)+1 to begin with
// That owuld have made it lot simpler

func rev_ar (ar []int) []int {
    ans:=[]int{}
    for i:=len(ar)-1;i>-1;i--{
        ans = append(ans,ar[i])
    }
    return ans
}



func conv_to_ar(l,n int) []int {
    // l = total numbers items in []int

    ans:=[]int{}


    q := n / 10
    d := n % 10
    ans = append(ans,d)

    for q > 0 {
        d = q % 10
        q = q / 10
        ans = append(ans,d)
    }

    for i:=0;len(ans)<l;i++ {
        ans = append(ans,0)
    }

    return rev_ar(ans)
}


func logbase10(n int) int {
    ans:= 0
    q := n
    for q>0 {
        q = q/10
        ans++
    }
    return ans
}

func addToArrayForm(num []int, k int) []int {
  

  //fmt.Println(logbase10(100))

  ans:= []int{}
  ar_of_k := conv_to_ar(len(num),k)
  //fmt.Println(ar_of_k)

  // if num is shorter than k, the I have to prepend num with 0s
  for len(num) < (logbase10(k)) {
    num= append([]int{0}, num...)
  }
  //fmt.Println(num)

  // fill up ans so that it has same length as num
  for i:=0;i<len(num);i++ {
    ans = append(ans,0)
  }


  // add nums and ar_of_k item by item and insert into "ans", from right to left
  carry_over := 0
  for i:=len(ar_of_k)-1;i>-1;i-- {
    if num[i] + ar_of_k[i]+carry_over <= 9 {
       ans[i] = num[i] + ar_of_k[i] + carry_over
       carry_over = 0
    } else { // more than 9
       ans[i] = (num[i] + ar_of_k[i] + carry_over) % 10
       carry_over = 1
    }
  }

  if carry_over == 1 {
    // we need to prepend with 1
    ans= append([]int{1}, ans...)
  }
  
  

  


  return ans


    
}
