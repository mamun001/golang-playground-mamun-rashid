

// leetcode easy #121 stock profit
// for loop
 
func maxProfit(prices []int) int {
    
    lowest_price := 10000
    //highest_price := 0
    could_sell := false
    highest_possible_profit := 0
    todays_profit := 0
    
    
    
    for day:=0;day<len(prices); day++ {
        if prices[day] < lowest_price {
            lowest_price = prices[day]
        } 
        
        
        
        if prices[day] > lowest_price {
            if day != 0 {
                could_sell = true
                todays_profit = prices[day] - lowest_price
                if todays_profit > highest_possible_profit {
                    highest_possible_profit = todays_profit       
                }            
            }  
        } 
    } 
    
    //fmt.Println(lowest_price)
    //fmt.Println(highest_price)
    //fmt.Println(highest_possible_profit)
    //fmt.Println(could_sell)
    
    
    if could_sell == true {
        return highest_possible_profit
    } else {
        return 0
    }
    
}
