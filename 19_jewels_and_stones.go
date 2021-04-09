func numJewelsInStones(jewels string, stones string) int {
    
    var count_of_matches int = 0
    var summy int = 0
    
    for i, c := range jewels {
        aORb := regexp.MustCompile(string(c))
        matches := aORb.FindAllStringIndex(stones, -1)
        count_of_matches = count_of_matches + len(matches)
        summy = summy + i
    
	}
    
    
    return count_of_matches
}
