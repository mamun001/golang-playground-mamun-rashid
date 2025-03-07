
// easy 3178, ELO 1256
// Find the Child Who Has the Ball


func numberOfChild(n int, k int) int {

    who := 0
    dir := "plus"

    for i:=1 ; i<=k; i++ {
        if dir == "plus" {
           who = who + 1 
        } else if dir == "minus" {
            who = who - 1
        }
        if who == n-1 {
            dir = "minus"
        } else if who == 0 {
            dir = "plus"
        }
    }

    return who
    
}
