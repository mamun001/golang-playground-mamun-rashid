// easy , 1784, ELO 1206
// most horrendous problem description
// I spent 45 minutes on this , only to fail
// This solution is golang version of java solution from youtube

func checkOnesSegment(s string) bool {
    
    for i:=1; i<len(s); i++ {
        if s[i-1] == '0' && s[i] == '1' {
            return false
        }
    }

    return true
    
}
