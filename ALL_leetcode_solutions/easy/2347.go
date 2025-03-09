// easy 2347, ELO 1242
// Best Poker Hand

func bestHand(ranks []int, suits []byte) string {

    if suits[0] == suits[1] && suits[1] == suits[2] && suits[2] == suits[3] && suits[3] == suits[4] {
        return "Flush"
    } 

    map_rank := make(map[int] int) 
        for i:=0; i<len(ranks);i++ {
            map_rank[ranks[i]]++
        }
    
    max_freq:=0 
    for _,v := range map_rank {
        if v > max_freq {
            max_freq = v
        }
    }
    if max_freq >=3 {
        return "Three of a Kind"
    }
    if max_freq == 2 {
        return "Pair"
    }

    max_rank := 0
    for i:=0;i<len(ranks);i++ {
        if ranks[i] > max_rank {
            max_rank = ranks[i]
        }
    }
    return "High Card"

    
}
