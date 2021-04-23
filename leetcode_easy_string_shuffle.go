


// leetcode easy level
// string

func restoreString(s string, indices []int) string {
    
    chars := make([]byte, len(indices))
    var result string
    
    for i:=0; i < len(indices); i++ {
        chars[indices[i]] = s[i]
        
    }
    
    result = string(chars[:])
    
    return result
    
}
