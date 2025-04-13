// easy 2404, ELO 1259
// most frequent even element

func mostFrequentEven(nums []int) int {

    freq_map := make(map[int]int)

    for i:=0;i<len(nums);i++ {
        if nums[i] % 2 == 0 {
            freq_map[nums[i]]++
        }
    }

    max:=-1
    winner:=-1

    for k,v := range freq_map {
        if v > max {
            max = v
            winner =k
        }
        if v == max && k < winner {
            winner =k
        }
    }

    return winner
    
}
