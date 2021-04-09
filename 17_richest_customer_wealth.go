import "fmt"

func maximumWealth(accounts [][]int) int {
    var sum int = 0
    var sum_array [100]int
    
    for i, row := range accounts {
        sum = 0
        for _, val := range row {
            sum = sum + val
        }
        sum_array[i] = sum
    }
    
    var most_wealth int = 0
    
    for j:=0; j < len(sum_array); j++ {
      if sum_array [j] > most_wealth {
          most_wealth = sum_array [j]
      }
    }
    
    return most_wealth
    
}
