func numberOfSteps(num int) int {
    if num == 0 {
        return 0
    } else if num == 1 {
            return 1
    } else if (num%2==0){
          // even
          return numberOfSteps(num /2) + 1
      } else {
          // odd
          return numberOfSteps((num -1) /2) +2
      }

}
