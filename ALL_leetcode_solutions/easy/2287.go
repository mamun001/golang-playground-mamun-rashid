// easy 2287, ELO 1300
// rearrange characters to make target string

func rearrangeCharacters(s string, target string) int {

    map_target := make(map[byte]int)
    for i:=0;i<len(target);i++ {
        map_target[target[i]]++
    }
    //fmt.Println(map_target)

    map_s := make(map[byte]int)
    for i:=0;i<len(s);i++ {
        map_s[s[i]]++
    }
    //fmt.Println(map_s)

    min_mult:=9999999
    for k,_ := range(map_target) {
        //fmt.Println(k,map_s[k])
        cur:=map_s[k]/map_target[k]
        if cur < min_mult {
            min_mult = cur
        }
    }
    


    return min_mult
    
}
