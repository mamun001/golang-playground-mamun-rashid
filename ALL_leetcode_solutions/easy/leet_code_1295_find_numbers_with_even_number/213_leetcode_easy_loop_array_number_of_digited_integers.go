// leetcode easy #1295 Find Numbers with Even Number of digits
// loop, array

/* My Test Cases
[1]
[20]
[1,2,3,4,5,6]
[20,20,30,40,50]
[1,20,3,40,400,500,6000]
[100000,99999,99998,99,1,50]
[100,2,46,9000,100,3]
*/


func num_of_digits(n int) int {
    s := strconv.Itoa(n) // convert to string
    return len(s)  
}


func findNumbers(nums []int) int {
    
    answer := 0
    
    //fmt.Println(num_of_digits(100000))
    
    for i:=0;i < len(nums); i++ {
        digits := num_of_digits(nums[i])
        if(digits%2==0){
          //fmt.Println(n,"is Even number")
          answer++
        } // if
    } // for
    
    
    
    return answer
    
}
