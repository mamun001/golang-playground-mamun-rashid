// 2806, easy, ELO 1215
// Account Balance After Rounded Purchase

func round(n int) int {
    ans:= 0

    if n % 10 <= 4 {
        return (n / 10) * 10
    } else {
        return ((n / 10) +1)  * 10 
    }
    return ans
}

func accountBalanceAfterPurchase(purchaseAmount int) int {

   //fmt.Println(round(15))




    return 100 - round(purchaseAmount)
    
}
