// easy 3184
// Count Pairs That Form
// ELO 1149 or so

func countCompleteDayPairs(hours []int) int {

    L := len(hours)
    ans := 0

    // brute force 
    // because we are doing brute force, each pair will come up twice
    for i:=0 ; i<L; i++ {
      for j:=0 ; j<L; j++ {
        if i != j && (hours[i]+hours[j]) % 24 == 0 {
            ans++
        }
      }
    }
    return ans / 2
}

