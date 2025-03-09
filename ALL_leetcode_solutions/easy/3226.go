// easy 3226, ELO 1247
// number of bit changes to make two integers equal

func minChanges(n int, k int) int {

    if n == k {
        return 0
    }

    if k > n {
        return -1
    }

    // n is guaranteed to larger than k at this point
    q_n := n // q for quotient
    q_k := k
    ans := 0

    for q_n > 0 {
       mod_n := q_n % 2
       mod_k := q_k % 2
       //fmt.Println(q_n,q_k)
       // if k's digit is 1 and n's digit on 0, game over 
       if mod_k == 1 && mod_n == 0 {
          //fmt.Println("k is 1 & n is 0, returning -1")
          return -1
       }
       if mod_k == 1 && mod_n == 1 {
          //fmt.Println("k is 1 & n is 1, doing nothing")
          ans = ans // do nothing
       }

       if mod_k == 0 && mod_n == 1 {
          //fmt.Println("k is 0 & n is 1, doing ans++")
          ans++ // it is still possible to make this happen
       }
       q_n = q_n / 2
       q_k = q_k / 2
    }



    return ans
    
}
