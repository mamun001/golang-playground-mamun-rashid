func CountDigits(i int) (count int) {
 	for i != 0 {

 		i /= 10
 		count = count + 1
 	}
 	return count
 }


func powInt(x, y int) int {
    return int(math.Pow(float64(x), float64(y)))
}

func subtractProductAndSum(n int) int {
    diff := 0   

    temp := n
    //total_digits := CountDigits(n)
    //tens := total_digits -1
    current_digit := 0
    previous_temp := n
    sum := 0
    mult := 1
    
  for  temp != 0 {
        previous_temp = temp
 		temp /= 10  // 234 to 23 to 2
        current_digit = previous_temp - (temp * 10)
        fmt.Println(current_digit)
        sum = sum + current_digit
        mult = mult * current_digit
      
 		
 	}

    fmt.Println(sum,mult)
  diff = mult - sum
  return diff
    
}
