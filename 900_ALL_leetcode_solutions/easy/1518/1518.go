// 70 pct acc rate

// easy level, 1518, Water Bottles

// DEC 26, 2024


import "fmt"

func numWaterBottles(numBottles int, numExchange int) int {

    bottles_drank := 0

    bottles_at_hand := numBottles

    fmt.Println ("start", bottles_at_hand,bottles_drank)

    for bottles_at_hand >= numExchange { // as long we have enough to exchange
        bottles_at_hand = bottles_at_hand - numExchange // drank as many as we can exchange
        
        bottles_drank = bottles_drank + numExchange
        fmt.Println ("drank 3")
        fmt.Println ("at hand:", bottles_at_hand,"drank:", bottles_drank)
        bottles_at_hand = bottles_at_hand + 1 // went to store and got one
        fmt.Println("went to store and got 1")
        fmt.Println ("at hand:", bottles_at_hand, "drank:", bottles_drank)
    }

    // drink the ones we cannot exchange
    fmt.Println ("drank leftovers")
    bottles_drank = bottles_drank + bottles_at_hand
    fmt.Println (bottles_at_hand,bottles_drank)


    return bottles_drank
    
}

