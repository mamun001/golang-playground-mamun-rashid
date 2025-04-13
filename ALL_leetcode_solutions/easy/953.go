// easy 953, ELO 1300
// verifying an alien dictionary


func index(b byte, order string) int{

    for i:=0;i<len(order);i++ {
       if b == order[i] {
        return i
       }
    }
    return -1
}




func is_correct(s1,s2,order string) bool {

    for i:=0;i<min(len(s1),len(s2));i++ {
        //fmt.Println(s1[i],s2[i])
        //fmt.Println(index(s1[i],order),index(s2[i],order)) 
        if index(s1[i],order) < index(s2[i],order) {      
              //fmt.Println("returning true")
              return true
        }
        if index(s1[i],order) > index(s2[i],order) {      
              //fmt.Println("returning false")
              return false
        }
        if index(s1[i],order) == index(s2[i],order) {      
              //fmt.Println("not deciding")
              continue
        }
    }
    // finished comparing up to length of the shorter one and still all equal
    if len(s1) > len(s2) {
       return false
    }
    if len(s2) > len(s1) {
       return true
    }

    fmt.Println("we should never be here")
    return true // we should never be here
}


func isAlienSorted(words []string, order string) bool {

    //fmt.Println(index('h',order))
    //fmt.Println(index('l',order))
    //fmt.Println(words[0],words[1])
    //fmt.Println(is_correct(words[0],words[1],order))

    for i:=0;i<len(words)-1;i++{
        if is_correct(words[i],words[i+1],order) == false {
            return false
        }
    }

    return true
    
}
