// easy 1848, Minimum Distance to the target
// ELO 1217


func abs (n int) int {
    if n < 0 {
        return n * -1
    } else {
        return n
    }
}

func getMinDistance(nums []int, target int, start int) int {

    min:= 10001
    this:=10001

    for i:=0;i<len(nums);i++ {
        if nums[i] == target {
            this = abs(i-start)
            if this < min {
                min = this
            }
        }
    }

    return min
    
}
