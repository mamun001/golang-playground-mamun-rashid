// EASY 1450 
// number of students
// 8th or so lowest ELO

func busyStudent(startTime []int, endTime []int, queryTime int) int {

    ans := 0

    for i:=0 ; i < len(startTime) ; i++ {
      if startTime[i] <= queryTime && endTime[i] >= queryTime {
        ans++
      }
    }
    
    return ans
}

