// easy 2399, ELO ELO 1244
// check distances between same letters

func checkDistances(s string, distance []int) bool {

    map1 := make(map[byte]int) // we will save index of the first occurance of each letter
    //fmt.Println(map1)

    for i:=0;i<len(s);i++ {
        _,exists := map1[s[i]]
        if exists == false {
            map1[s[i]]=i
        } else {
            dist := i-map1[s[i]]-1 // dumb definition of "distance" given by problem description, but ok
            //fmt.Println(s[i],dist)
            index:=s[i]-97 // index of the given letter in the distances array
            //fmt.Println(index)
            if dist != distance[index] { // check for the correct value
                return false
            }
        }
    }
    //fmt.Println(map1)



    return true
    
}
