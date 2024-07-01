
// compare two slices item by item
// code by Mamun Rashid
// tested some

func compare_slices (s1 []int, s2 []int) bool {
    answer := true // we start assuming its true and make it false down the road if it false
    l1 := len(s1)
    l2 := len(s2)
    
    if l1 != l2 {
        return false // quick turn around
    } else {
        for i:=0; i < l1; i++ {
            if s1[i] == s2[i] {
                continue
            } else {
                answer = false
                return answer
                break
            }
        }
    }
    
    return answer
}
