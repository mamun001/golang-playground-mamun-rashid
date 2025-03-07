/*

Leetcode level = easy 1588
Leetcode Problem  Title:  sum-of-all-odd-length-subarrays
Key CS topics: nested loop, math

My test cases
[1]
[1,2]
[1,2,3]
[1,2,3,4]
[1,2,3,4,5]
[1,2,3,4,5,6]
[1,2,3,4,5,6,7]
[1,2,3,4,5,6,7,8]
[1,2,3,4,5,6,7,8,9]
*/


func sumOddLengthSubarrays(arr []int) int {
    
    answer := 0
    
    L:=len(arr)
    
    max_odd := 0  // maximum of odd-length subarray possible
                  // e.g. if arr is of length 6, max_odd is 5
    if(L%2==0){
        //even length array
        max_odd = L-1
    }else{
        //odd length array
        max_odd = L
    }
    
    
    //fmt.Println("max_odd_possible:", max_odd)
    
    slice_of_indexes := make([]int, 0)  // VERY LONG INDEXES THAT WILL BE SUMMED UP IN THE END
    
    for i:=1; i <= max_odd; i=i+2 { // i=how many we sub-arraying at a time
        //fmt.Println("i=",i)
        for j:=0; j <= L-i; j++ { // indexes that fall in -sliding indexes
            //fmt.Println("j=",j)
            for k:=j; k < j+i; k++ {
               //fmt.Println("k=",k)
               slice_of_indexes = append (slice_of_indexes, k)  // k is the index we finally care about
               //fmt.Println("i,j,k:", i,j,k)
            } // k
          } // j
        //fmt.Println(slice_of_indexes)      
    } // i
    
    
    // now sum up values of LONG list of indexes over and over again
    
    for m:=0;m<len(slice_of_indexes);m++ {
        answer = answer + arr[slice_of_indexes[m]]
    }
    
    return answer
    
}
