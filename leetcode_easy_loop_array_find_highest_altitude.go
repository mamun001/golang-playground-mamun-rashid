/*
leetcode easy #1732 Gind the Highest Altitude
loop, array

my test data

[-5,1,5,0,-7]
[-4,-3,-2,-1,4,3,2]
[-1]
[2,-5]
[2,2,2,2,2]
[-1,-2,-3]
[-10,1,1,1,1,1,1,1]

*/



func largestAltitude(gain []int) int {
    
    running_height := 0
    max_height_so_far := 0
    
    for i := 0; i < len(gain) ; i++ {
      running_height = running_height + gain[i]
    
      if running_height > max_height_so_far {
         max_height_so_far = running_height
      }
    }
    
    return max_height_so_far
}
