import "fmt"

// take one from first half, then one from second half and the repeat
// input = input array and half of number of elements in array.
// [6,5,1,3,4,7]
// 3  
// output =
// [6,3,5,4,1,7]


func shuffle(nums []int, n int) []int {
    firsthalf := make([]int, n)
    secondhalf := make([]int, n)
    final := make([]int, n*2)
    
    for i:=0;i<n;i++ {
        firsthalf[i]=nums[i]
    }
    
    j:=0
    for i:=n;i<n*2;i++ {
        secondhalf[j]=nums[i]
        j++
    }
    
    index_at_the_halves:=0
    for i:=0;i<n*2;i++ {
          // even then odd
        if(i%2==0){
          final[i]=firsthalf[index_at_the_halves]
        }else{
          final[i]=secondhalf[index_at_the_halves]
          index_at_the_halves++
        }
        
        
     }
    
    
    return final
    
}
