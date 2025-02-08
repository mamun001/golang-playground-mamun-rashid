// easy 1491 , ELO 1201
// Average Salary Excluding the

import "sort"

func average(salary []int) float64 {

  sort.Ints(salary)

  sum := 0.0

  for i:=1;i<len(salary)-1;i++ {
    sum = sum + float64(salary[i])
  }

  return sum/float64(len(salary)-2)

    
}

