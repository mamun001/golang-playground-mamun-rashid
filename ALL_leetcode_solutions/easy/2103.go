// easy 2103 , ELO 1258
// rings and rods
// should have used a custom func with sort.Slice func
// that would have taken lot shorter. I spend 90+ minutes on this

func is_in(ar []byte, b byte) bool {
    for i:=0;i<len(ar);i++ {
        if ar[i] == b {
            return true
        }
    }
    return false
}


func is_all_3(ar []byte) bool {
    if len(ar) < 3 {
        return false
    }
    if is_in(ar,'B') && is_in(ar,'G') && is_in(ar,'R') {
        return true
    }
    return false
}

func countPoints(rings string) int {

    // 0 => 48
    // B/G/R 66,71,82

    map_rings:=make(map[byte][]byte)

    for i:=0;i<len(rings);i++ {
        if i % 2 == 1 {
            map_rings[rings[i]]=append(map_rings[rings[i]],rings[i-1])
        }
    }
    //fmt.Println(map_rings)

    ans:=0
    for _,v := range map_rings {
        if is_all_3(v) {
            ans++
        }
    }


    return ans
    
}
