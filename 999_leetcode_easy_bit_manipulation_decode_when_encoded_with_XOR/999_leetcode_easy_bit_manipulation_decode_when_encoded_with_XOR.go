


func decode(encoded []int, first int) []int {
    
    result := make([]int, 0)
    
    result = append (result,first)
    
    next := 0
    
    for i:=0; i<len(encoded); i++ {
        next = result[i] ^ encoded[i]
        result = append (result, next)
    }
    
    return result
    
}
