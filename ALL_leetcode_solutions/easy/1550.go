// easy 1550 ELO 1221 
// three consecutive odds

func threeConsecutiveOdds(arr []int) bool {

    if len(arr) < 3 {
        return false
    }

    ans := false

    for i:=0;i<len(arr)-2;i++ {
        if arr[i] % 2 == 1 && arr[i+1] % 2 == 1 && arr[i+2] % 2 == 1 {
            return true
        }
    }


    return ans
    
}
