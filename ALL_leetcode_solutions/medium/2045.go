// 2045 , medium, ELO 1355
// 3/31 needed help
// optimal partition of string

func partitionString(s string) int {

    map_bytes := make(map[byte]int)

    substr_count := 0

    for i:=0;i<len(s);i++ {
        if map_bytes[s[i]] == 1 {
            //fmt.Println(i,s[i])
            //fmt.Println("increasing count")
            substr_count++
            for k,_ := range map_bytes {
                map_bytes[k] = 0
            }
            map_bytes[s[i]] = 1
            //fmt.Println(map_bytes)
        } else {
            map_bytes[s[i]] = 1
        }
    }

    //fmt.Println(map_byte_freq)


    return substr_count+1
    
}
