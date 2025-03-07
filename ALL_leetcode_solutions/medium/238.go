// Medium 238, no ELO
// part of neetcode 250
// Had to get the algorithm from Youtube

func productExceptSelf(nums []int) []int {

    L :=[]int{} // "prefix" array
    R :=[]int{} // "postfix" array
    ans :=[]int{} // answer
    n := len(nums) 
    

    // prefill the arrays, because that makes rest of code much easier to write
    for i:=0;i<n;i++{
        L = append(L,1)
        R = append(R,1)
        ans = append(ans,1)
    }

    // build the prefix array
    L[0] = 1
    for i:=1;i<n;i++{
        L[i] = L[i-1] * nums[i-1]
    }

    // build the postfix array
    R[n-1] = 1
    for i:=n-2;i>-1;i--{
      R[i] = R[i+1] * nums[i+1]  
    }

    // simple to fill in answer once you know
    for i:=0;i<n;i++{
        ans[i] = L[i] * R[i]
    }


    //fmt.Println(L)
    //fmt.Println(R)
    //fmt.Println(ans)


    return ans
    
}
