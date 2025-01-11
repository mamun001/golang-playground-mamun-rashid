// easy 1475
// final prices with a special discount in a shop
// 83 pct acc rate

// using monolithic stack did not make sense for simple one with only 50 items in the array

// my algorithm: < 1ms beats "100" pct
// RAM: beats 12%, so not great, but ok



func finalPrices(prices []int) []int {

    ans := []int{}
    this_price := 0
    found_cheaper := false

    for i:=0 ; i < len(prices) ; i++ {
        found_cheaper = false
        for j:=i+1 ; j < len(prices) ; j++ {
            if prices[j] <= prices[i] { // found a price that is less than current AFTER the current item
                this_price = prices[i] - prices[j]
                ans = append(ans,this_price)
                j = len(prices) // break from inner loop
                found_cheaper = true
            } 
        } // for inner loop ends
        if found_cheaper == false { // no cheaper price was found here onwards
            ans = append(ans,prices[i]) // add undiscouted price to ans
        }
    } // for outer

    return ans
    
}

