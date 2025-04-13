// easy 2283 , ELO 1253
// check if NUmber Has Equal Digit Count and Digit Value

func to_byte(i int) byte {

    switch i {
        case 0: return '0'
        case 1: return '1'
        case 2: return '2'
        case 3: return '3'
        case 4: return '4'
        case 5: return '5'
        case 6: return '6'
        case 7: return '7'
        case 8: return '8'
        case 9: return '9'
    }
    return '1'   
}

func to_int(b byte) int {

    switch b {
        case '0': return 0
        case '1': return 1
        case '2': return 2
        case '3': return 3
        case '4': return 4
        case '5': return 5
        case '6': return 6
        case '7': return 7
        case '8': return 8
        case '9': return 9
    }
    return 0   
}



func digitCount(num string) bool {

    // how many times has a digit shown up?
    map_freq:= make(map[byte]int)
    for i:=0;i<len(num);i++ {
        map_freq[byte(num[i])]++
    }

    //fmt.Println(map_freq)
    // 0 = 48 (byte)
    // 9 = 57 (byte)
    
    
    for i:=0;i<len(num);i++ {
        //fmt.Println(i)
        //fmt.Println(to_int(num[i]))
        //fmt.Println(map_freq[to_byte(i)])
        // ACTUAL CHECK
        if to_int(num[i]) != map_freq[to_byte(i)] {
            return false
        }
    }

    return true
    
}
