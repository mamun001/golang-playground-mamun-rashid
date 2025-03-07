// MEDIUM , 122, no ELO
// part of neetcode 250

// Had to lookup the algorithm in chatgpt
// once found, super simple to implement

func maxProfit(prices []int) int {

    profit := 0

    for i:=0;i<len(prices)-1;i++ {
        if prices[i+1] > prices[i] {
            profit = profit + prices[i+1] - prices[i]
        }
    }
    
    return profit
}
