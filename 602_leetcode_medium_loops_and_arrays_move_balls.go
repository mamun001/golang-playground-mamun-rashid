
// Idea: answer of each entry is just the sum of distances of all the existing balls

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}



func minOperations(boxes string) []int {
    
    answer := make([]int, len(boxes))
    
    position_of_balls := make([]int, 0)
    
    for pos, char := range boxes {
        if char == '1' {
            position_of_balls = append(position_of_balls, pos)
        }
    }
    
    for i:=0; i < len(position_of_balls); i++ {
        fmt.Println(position_of_balls[i])
    }
    
    for i:=0; i < len(answer); i++ {
        sum_of_distances := 0
        for j:=0 ; j < len(position_of_balls) ; j++ {
          sum_of_distances = sum_of_distances + Abs(i-position_of_balls[j])
        }
        answer[i] = sum_of_distances
    }
    
    return answer
    
}
